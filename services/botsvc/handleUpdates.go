package botsvc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	c "github.com/gpng/delivery-bot-api/constants"
	"github.com/gpng/delivery-bot-api/models"
	"github.com/gpng/delivery-bot-api/utils/deliveryslots"
	u "github.com/gpng/delivery-bot-api/utils/utils"
	"github.com/jinzhu/gorm"
)

func (s *Service) handleUpdates() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.render.Respond(w, r, s.render.Message(true, "ok"))
		update := &models.TelegramUpdate{}
		if err := json.NewDecoder(r.Body).Decode(update); err != nil {
			u.LogError(err)
			return
		}

		chatID := update.Message.Chat.ID
		text := update.Message.Text
		split := strings.Split(text, " ")
		switch split[0] {
		case "/start":
			s.handleStart(chatID)
		case "/help":
			s.handleHelp(chatID)
		case "/postalcode":
			if len(split) < 2 {
				s.bot.SendMessage(chatID, "Invalid postal code")
			} else {
				s.handlePostalcode(chatID, split[1])
			}
		case "/check":
			s.handleCheck(chatID)
		case "/pause":
			s.handlePause(chatID)
		case "/unpause":
			s.handleUnpause(chatID)
		}
	}
}

func (s *Service) handleStart(chatID int64) {
	s.bot.SendMessage(chatID, "Welcome to SG Delivery Slots Bot. Use \"/postalcode <postal code>\" (e.g. \"/postalcode 123456\") to set your postal code and start tracking. Use /pause to pause updates and /unpause to restart, and /check to manually trigger an update.")
}

func (s *Service) handleHelp(chatID int64) {
	s.bot.SendMessage(chatID, "Use \"/postalcode <postal code>\" (e.g. \"/postalcode 123456\") to set your postal code and start tracking. Use /pause to pause updates and /unpause to restart, and /check to manually trigger an update.")
}

func (s *Service) handlePostalcode(chatID int64, postcode string) {
	match, _ := regexp.MatchString("^\\d{6}$", postcode)
	if !match {
		s.bot.SendMessage(chatID, "Invalid postal code, must consists of 6 digits only (e.g. /postalcode 123456)")
		return
	}

	// check if postcode already exists
	existingPostcode, err := models.GetPostcodeByChatID(s.db, chatID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return
	}
	// no existing postcode, create
	if gorm.IsRecordNotFoundError(err) {
		newPostcode := models.Postcode{
			PostcodeString: postcode,
			ChatID:         chatID,
			Status:         c.StatusActive,
		}
		if err := newPostcode.Create(s.db); err != nil {
			s.bot.SendMessage(chatID, "Failed to update postcode, please contact dev :(")
			return
		}

		s.bot.SendMessage(chatID, fmt.Sprintf("Success! Your postcode has been set to %s and tracking has started! Use \"/pause\" to pause tracking and \"/unpause\" to start tracking again.", postcode))
		deliveryslots.CheckAll(s.db, s.bot, []int64{chatID}, postcode, true)
		return
	}

	// otherwise update
	if err := models.UpdatePostcode(s.db, existingPostcode.ID, postcode); err != nil {
		s.bot.SendMessage(chatID, "Failed to update postcode, please contact dev :(")
		return
	}

	s.bot.SendMessage(chatID, fmt.Sprintf("Success! Your postcode has been set to %s.", postcode))

	deliveryslots.CheckAll(s.db, s.bot, []int64{chatID}, postcode, true)
}

func (s *Service) handleCheck(chatID int64) {
	// check if postcode already exists
	existingPostcode, err := models.GetPostcodeByChatID(s.db, chatID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return
	}
	if gorm.IsRecordNotFoundError(err) || existingPostcode.PostcodeString == "" {
		s.bot.SendMessage(chatID, "You have not set a postal code yet. Use \"/postalcode <postal code>\" (e.g. \"/postalcode 123456\" ) to set a postal code and start tracking.")
		return
	}

	deliveryslots.CheckAll(s.db, s.bot, []int64{chatID}, existingPostcode.PostcodeString, true)
}

func (s *Service) handlePause(chatID int64) {
	// check if postcode already exists
	existingPostcode, err := models.GetPostcodeByChatID(s.db, chatID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return
	}
	if gorm.IsRecordNotFoundError(err) || existingPostcode.PostcodeString == "" {
		s.bot.SendMessage(chatID, "You have not set a postal code yet. Use \"/postalcode <postal code>\" (e.g. \"/postalcode 123456\" ) to set a postal code and start tracking.")
		return
	}

	if err := models.UpdateStatus(s.db, existingPostcode.ID, c.StatusPaused); err != nil {
		s.bot.SendMessage(chatID, "Failed to update tracking status, please contact dev :(")
		return
	}
	s.bot.SendMessage(chatID, "Tracking stopped! Use /unpause to start tracking again.")
}

func (s *Service) handleUnpause(chatID int64) {
	// check if postcode already exists
	existingPostcode, err := models.GetPostcodeByChatID(s.db, chatID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return
	}
	if gorm.IsRecordNotFoundError(err) || existingPostcode.PostcodeString == "" {
		s.bot.SendMessage(chatID, "You have not set a postal code yet. Use \"/postalcode <postal code>\" (e.g. \"/postalcode 123456\" ) to set a postal code and start tracking.")
		return
	}

	if err := models.UpdateStatus(s.db, existingPostcode.ID, c.StatusActive); err != nil {
		s.bot.SendMessage(chatID, "Failed to update tracking status, please contact dev :(")
		return
	}
	s.bot.SendMessage(chatID, "Tracking started! Use /pause to pause tracking.")
	deliveryslots.CheckAll(s.db, s.bot, []int64{chatID}, existingPostcode.PostcodeString, true)
}
