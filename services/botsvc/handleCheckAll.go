package botsvc

import (
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
		postcodeMap := map[string][]int64{}
		// go by distinct postcodes

		for _, postcode := range postcodes {
			if val, ok := postcodeMap[postcode.PostcodeString]; ok {
				postcodeMap[postcode.PostcodeString] = []int64{postcode.ChatID}
			} else {
				postcodeMap[postcode.PostcodeString] = append(val, postcode.ChatID)
			}
		}
		for postcode, chatIDs := range postcodeMap {
			deliveryslots.CheckAll(s.db, s.bot, chatIDs, postcode, false)
		}
	}
}
