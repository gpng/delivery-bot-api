package deliveryslots

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gpng/delivery-bot-api/connections/telegram"
	c "github.com/gpng/delivery-bot-api/constants"
	u "github.com/gpng/delivery-bot-api/utils/utils"
	"github.com/jinzhu/gorm"
)

type coldstorageEarliestResponse struct {
	Date  string `json:"date"`
	Label string `json:"label"`
}

type coldstorageResponse struct {
	Earliest json.RawMessage `json:"earliest"`
}

func checkColdstorage(db *gorm.DB, bot *telegram.Bot, chatIDs []int64, postcode string, negativeResponse bool) {
	message := ""
	available := false

	url := "https://coldstorage.com.sg/checkout/cart/checkdelivery"

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormField("postal_code")
	if err != nil {
		u.LogError(err)
		return
	}
	io.Copy(fw, strings.NewReader(postcode))
	w.Close()

	resp, err := http.Post(url, w.FormDataContentType(), &b)
	if err != nil {
		u.LogError(err)
		return
	}

	decoded := &coldstorageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		u.LogError(err)
		return
	}
	// 2 cos empty message is "[]"
	if len(decoded.Earliest) > 2 {
		earliest := coldstorageEarliestResponse{}
		err := json.Unmarshal(decoded.Earliest, &earliest)
		available = true
		if err != nil {
			message = "Cold Storage slot available! Go to https://coldstorage.com.sg/checkout/cart"
		} else {
			message = fmt.Sprintf("Cold Storage available earliest on %s from %s! Go to https://coldstorage.com.sg/checkout/cart", earliest.Date, earliest.Label)
		}
	} else if negativeResponse {
		message = "No Cold Storage slots available :("
	}

	notify(db, bot, negativeResponse, available, message, chatIDs, c.ServiceColdstorage)
}
