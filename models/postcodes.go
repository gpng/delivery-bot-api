package models

import (
	c "github.com/gpng/delivery-bot-api/constants"
	u "github.com/gpng/delivery-bot-api/utils/utils"
	"github.com/jinzhu/gorm"
)

// Postcode model
type Postcode struct {
	Model
	PostcodeString string `json:"postcodeString"`
	Postcode       int    `json:"postcode"`
	ChatID         int64  `json:"chatId"`
	Status         int    `json:"status"`
}

// GetPostcodeByChatID from db
func GetPostcodeByChatID(db *gorm.DB, chatID int64) (Postcode, error) {
	postcode := Postcode{}
	err := db.Where("chat_id = ?", chatID).First(&postcode).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		u.LogError(err)
	}
	return postcode, err
}

// Create new postcode
func (postcode *Postcode) Create(db *gorm.DB) error {
	err := db.Create(postcode).Error
	if err != nil {
		u.LogError(err)
	}
	return err
}

// UpdatePostcode by chat id
func UpdatePostcode(db *gorm.DB, id string, postcode string) error {
	postcodeModel := Postcode{}
	postcodeModel.ID = id
	err := db.Model(postcodeModel).Update("postcode_string", postcode).Error
	if err != nil {
		u.LogError(err)
	}
	return err
}

// UpdateStatus by chat id
func UpdateStatus(db *gorm.DB, id string, status int) error {
	postcodeModel := Postcode{}
	postcodeModel.ID = id
	err := db.Model(postcodeModel).Update("status", status).Error
	if err != nil {
		u.LogError(err)
	}
	return err
}

// GetActivePostcodes where status is active and postcode is legit
func GetActivePostcodes(db *gorm.DB) ([]Postcode, error) {
	postcodes := []Postcode{}
	err := db.Where("status = ?", c.StatusActive).Find(&postcodes).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		u.LogError(err)
	}
	return postcodes, err
}
