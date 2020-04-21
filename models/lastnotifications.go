package models

import (
	u "github.com/gpng/delivery-bot-api/utils/utils"
	"github.com/jinzhu/gorm"
)

// LastNotification model
type LastNotification struct {
	Model
	ChatID  int64  `json:"chatId"`
	Service string `json:"service"`
}

// DeleteLastNotifications clears all by chat id
func DeleteLastNotifications(db *gorm.DB, chatID int64) error {
	err := db.Unscoped().Where("chat_id = ?", chatID).Delete(LastNotification{}).Error
	if err != nil {
		u.LogError(err)
	}
	return err
}

// DeleteLastNotificationsByService clears all by chat id for specific service
func DeleteLastNotificationsByService(db *gorm.DB, chatID int64, service string) error {
	err := db.Unscoped().Where("chat_id = ?", chatID).Where("service = ?", service).Delete(LastNotification{}).Error
	if err != nil {
		u.LogError(err)
	}
	return err
}

// Create last notification
func (lastNotification *LastNotification) Create(db *gorm.DB) error {
	err := db.Create(lastNotification).Error
	if err != nil {
		u.LogError(err)
	}
	return err
}

// GetLastNotification by chat id and service
func GetLastNotification(db *gorm.DB, chatIDs []int64, service string) ([]LastNotification, error) {
	lastNotifications := []LastNotification{}
	err := db.Where("chat_id IN (?)", chatIDs).Where("service = ?", service).Find(&lastNotifications).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		u.LogError(err)
	}
	return lastNotifications, err
}
