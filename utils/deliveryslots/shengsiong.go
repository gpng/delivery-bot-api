package deliveryslots

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gpng/delivery-bot-api/connections/telegram"
	c "github.com/gpng/delivery-bot-api/constants"
	u "github.com/gpng/delivery-bot-api/utils/utils"
	"github.com/jinzhu/gorm"
)

type shengsiongResponse struct {
	Response string `json:"response"`
}

func checkShengsiong(db *gorm.DB, bot *telegram.Bot, chatIDs []int64, postcode string, negativeResponse bool) {
	message := ""
	available := false

	url := "https://www.allforyou.sg/Common/pinCodeSearch"

	data := []byte(fmt.Sprintf(`{"pinStatus": 1, "code": "%s"}`, postcode))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		u.LogError(err)
		return
	}

	decoded := &shengsiongResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		u.LogError(err)
		return
	}

	if decoded.Response != "Failed" {
		available = true
		message = "Sheng Siong slot available! Go to https://www.allforyou.sg/cart"
	} else if negativeResponse {
		message = "No Sheng Siong slots available :("
	}

	notify(db, bot, negativeResponse, available, message, chatIDs, c.ServiceShengsiong)
}
