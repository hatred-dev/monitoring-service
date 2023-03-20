package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"monitoring-service/src/configuration"
	sm "monitoring-service/src/services/models"
	"net/http"
)

func SendTelegramNotification(text string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", configuration.TGConfig.BotToken)
	body, err := json.Marshal(sm.TelegramMessage{
		ParseMode: "MarkdownV2",
		Text:      text,
		ChatId:    configuration.TGConfig.ChatId,
	})
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	response, _ := io.ReadAll(resp.Body)
	fmt.Println(string(response))
}

func SendTimestamp() {}
