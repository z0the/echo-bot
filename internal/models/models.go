package models

type Chat struct {
	ID int `json:"id"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Update struct {
	ID      int     `json:"update_id"`
	Message Message `json:"message"`
}

type GetUpdatesResult struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}
