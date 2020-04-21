package deliveryslots

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/gpng/delivery-bot-api/connections/telegram"
	c "github.com/gpng/delivery-bot-api/constants"
	u "github.com/gpng/delivery-bot-api/utils/utils"
	"github.com/jinzhu/gorm"
)

type giantResponseEarliest struct {
	Date  string `json:"date"`
	Label string `json:"label"`
}

type giantResponse struct {
	Earliest json.RawMessage `json:"earliest"`
}

func checkGiant(db *gorm.DB, bot *telegram.Bot, chatIDs []int64, postcode int, negativeResponse bool) {
	message := ""
	available := false

	url := "https://giant.sg/checkout/cart/checkdelivery"

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormField("postal_code")
	if err != nil {
		u.LogError(err)
		return
	}
	io.Copy(fw, strings.NewReader(strconv.Itoa(postcode)))
	w.Close()

	resp, err := http.Post(url, w.FormDataContentType(), &b)
	if err != nil {
		u.LogError(err)
		return
	}

	decoded := &giantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		u.LogError(err)
		return
	}
	// 2 cos empty message is "[]"
	if len(decoded.Earliest) > 2 {
		available = true
		earliest := giantResponseEarliest{}
		err := json.Unmarshal(decoded.Earliest, &earliest)
		if err != nil {
			message = "Giant slot available! Go to https://giant.sg/checkout/cart"
		} else {
			message = fmt.Sprintf("Giant slot available earliest on %s from %s! Go to https://giant.sg/checkout/cart", earliest.Date, earliest.Label)
		}
	} else if negativeResponse {
		message = "No Giant slots available :("
	}

	notify(db, bot, negativeResponse, available, message, chatIDs, c.ServiceGiant)
}
