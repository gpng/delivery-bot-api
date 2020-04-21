package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gpng/delivery-bot-api/config"
	u "github.com/gpng/delivery-bot-api/utils/utils"
)

// Bot with all methods
type Bot struct {
	BotAPI tgbotapi.BotAPI
}

// New db connection and trigger migrations
func New(conf *config.Config) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(conf.Bot.Token)
	if err != nil {
		u.LogError(err)
		return nil, err
	}
	// connection string
	return &Bot{*bot}, nil
}

// SendMessage util
func (bot *Bot) SendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.DisableWebPagePreview = true
	bot.BotAPI.Send(msg)
}
