package domain

type MessageRequest struct {
	ChatID  string `json:"chat_id"`
	Message string `json:"message"`
}
