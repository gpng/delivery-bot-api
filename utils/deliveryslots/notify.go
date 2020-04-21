package deliveryslots

import (
	"github.com/gpng/delivery-bot-api/connections/telegram"
	"github.com/gpng/delivery-bot-api/models"
	u "github.com/gpng/delivery-bot-api/utils/utils"
	"github.com/jinzhu/gorm"
)

func notify(db *gorm.DB, bot *telegram.Bot, negativeResponse bool, available bool, message string, chatIDs []int64, service string) {
	lastNotifications, err := models.GetLastNotification(db, chatIDs, service)
	if err != nil {
		lastNotifications = []models.LastNotification{}
	}
	lastNotifiedMap := map[int64]bool{}
	for _, v := range lastNotifications {
		lastNotifiedMap[v.ChatID] = true
	}

	for _, chatID := range chatIDs {
		_, notified := lastNotifiedMap[chatID]
		if message != "" && (negativeResponse || !notified) {
			bot.SendMessage(chatID, message)
		}

		if available && !notified {
			lastNotification := models.LastNotification{
				ChatID:  chatID,
				Service: service,
			}
			if err := lastNotification.Create(db); err != nil {
				u.LogError(err)
			}
		} else if !available && notified {
			if err := models.DeleteLastNotificationsByService(db, chatID, service); err != nil {
				u.LogError(err)
			}
		}
	}
}
