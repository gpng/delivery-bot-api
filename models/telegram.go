package models

// Chat model
type Chat struct {
	ID int64 `json:"id"`
}

// Message model
type Message struct {
	MessageID int    `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Text      string `json:"text"`
}

// TelegramUpdate model
type TelegramUpdate struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}
