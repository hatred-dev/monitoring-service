package api

type TelegramMessage struct {
	ParseMode string `json:"parse_mode"`
	Text      string `json:"text"`
	ChatId    string `json:"chat_id"`
}

type UptimeNotification struct {
	ProjectKey string `json:"project_key"`
	Service    string `json:"service"`
	Status     string `json:"status"`
	Timestamp  string `json:"timestamp"`
}
