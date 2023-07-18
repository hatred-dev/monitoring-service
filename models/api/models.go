package api

type Ip struct {
	Ip     string
	Active bool
}

type Service struct {
	Name   string
	Url    string
	Active bool
}

type ProjectFull struct {
	ProjectName string
	Ips         []Ip
	Services    []Service
}

type TelegramMessage struct {
	ParseMode string `json:"parse_mode"`
	Text      string `json:"text"`
	ChatId    string `json:"chat_id"`
}

type ChannelStorage struct {
	HealthcheckChan chan bool
	PingChan        chan bool
}

type UptimeNotification struct {
	ProjectKey string `json:"project_key"`
	Service    string `json:"service"`
	Status     string `json:"status"`
	Timestamp  string `json:"timestamp"`
}
