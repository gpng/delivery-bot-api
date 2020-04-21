package botsvc

import (
	"fmt"
	"net/http"

	"github.com/gpng/delivery-bot-api/models"
	"github.com/gpng/delivery-bot-api/utils/deliveryslots"
	"github.com/jinzhu/gorm"
)

func (s *Service) handleCheckAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.render.Respond(w, r, s.render.Message(true, "ok"))

		postcodes, err := models.GetActivePostcodes(s.db)
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return
		}

		postcodeMap := map[int][]int64{}
		// go by distinct postcodes

		for _, postcode := range postcodes {
			if val, ok := postcodeMap[postcode.Postcode]; ok {
				postcodeMap[postcode.Postcode] = []int64{postcode.ChatID}
			} else {
				postcodeMap[postcode.Postcode] = append(val, postcode.ChatID)
			}
		}
		for postcode, chatIDs := range postcodeMap {
			deliveryslots.CheckAll(s.db, s.bot, chatIDs, postcode, false)
		}
	}
}

func (s *Service) handleCheckPostalCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.render.Respond(w, r, s.render.Message(true, "ok"))

		invalidPostcodes, err := models.GetInvalidPostcodes(s.db)
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return
		}

		for _, postcode := range invalidPostcodes {
			s.bot.SendMessage(postcode.ChatID, fmt.Sprintf("Your postcode %d is invalid, please set it again using /postcode <postal_code>", postcode.Postcode))
		}
	}
}
