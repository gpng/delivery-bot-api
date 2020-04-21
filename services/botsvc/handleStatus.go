package botsvc

import (
	"net/http"

	u "github.com/gpng/delivery-bot-api/utils/utils"
)

type statusResponse struct {
	Service string `json:"service"`
	Version int    `json:"version"`
}

// handleStatus returns the current api version
func (s *Service) handleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := statusResponse{
			Service: "BotSvc",
			Version: 1,
		}
		u.Respond(w, u.DataMessage(status, true, "API responding"))
	}
}
