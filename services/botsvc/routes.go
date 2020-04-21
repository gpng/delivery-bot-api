package botsvc

import (
	"github.com/go-chi/chi"
)

// Routes for app
func (s *Service) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", s.handleStatus())

	router.Post("/updates", s.handleUpdates())
	router.Get("/checkAll", s.handleCheckAll())
	router.Get("/checkPostcodes", s.handleCheckPostalCode())

	return router
}
