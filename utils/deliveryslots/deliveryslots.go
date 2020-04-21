package deliveryslots

import (
	"github.com/gpng/delivery-bot-api/connections/telegram"
	"github.com/jinzhu/gorm"
)

// CheckAll available services
func CheckAll(db *gorm.DB, bot *telegram.Bot, chatIDs []int64, postcode string, negativeResponse bool) {
	go checkColdstorage(db, bot, chatIDs, postcode, negativeResponse)
	go checkFairprice(db, bot, chatIDs, postcode, negativeResponse)
	go checkGiant(db, bot, chatIDs, postcode, negativeResponse)
	go checkShengsiong(db, bot, chatIDs, postcode, negativeResponse)
}
