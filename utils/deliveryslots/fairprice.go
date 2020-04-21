package deliveryslots

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gpng/delivery-bot-api/connections/telegram"
	c "github.com/gpng/delivery-bot-api/constants"
	u "github.com/gpng/delivery-bot-api/utils/utils"
	"github.com/jinzhu/gorm"
)

type fairpriceResponseData struct {
	Available bool `json:"available"`
}

type fairpriceResponse struct {
	Data fairpriceResponseData `json:"data"`
}

func checkFairprice(db *gorm.DB, bot *telegram.Bot, chatIDs []int64, postcode int, negativeResponse bool) {
	message := ""
	available := false

	url, err := url.Parse("https://website-api.omni.fairprice.com.sg/api/slot-availability")
	if err != nil {
		u.LogError(err)
		return
	}
	q := url.Query()
	q.Add("address[pincode]", strconv.Itoa(postcode))
	q.Add("storeId", "165")
	url.RawQuery = q.Encode()
	resp, err := http.Get(url.String())
	if err != nil {
		u.LogError(err)
		return
	}
	defer resp.Body.Close()

	decoded := &fairpriceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		u.LogError(err)
		return
	}
	if decoded.Data.Available {
		available = true
		message = "Fairprice slot available! Go to https://www.fairprice.com.sg/cart"
	} else if negativeResponse {
		message = "No Fairprice slots available :("
	}

	notify(db, bot, negativeResponse, available, message, chatIDs, c.ServiceFairprice)
}
