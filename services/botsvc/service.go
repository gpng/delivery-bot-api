package botsvc

import (
	"github.com/gpng/delivery-bot-api/connections/telegram"
	"github.com/gpng/delivery-bot-api/utils/render"
	vr "github.com/gpng/delivery-bot-api/utils/validator"
	"github.com/jinzhu/gorm"
)

// Service struct
type Service struct {
	db        *gorm.DB
	validator *vr.Validator
	render    *render.Render
	bot       *telegram.Bot
}

// New service
func New(
	db *gorm.DB,
	validator *vr.Validator,
	render *render.Render,
	bot *telegram.Bot,
) *Service {
	return &Service{
		db,
		validator,
		render,
		bot,
	}
}
