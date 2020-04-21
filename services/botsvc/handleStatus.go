package botsvc

import (
	"net/http"
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
		s.render.Respond(w, r, s.render.DataMessage(status, true, "API responding"))
	}
}
