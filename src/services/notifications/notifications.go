package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"monitoring-service/src/configuration"
	sm "monitoring-service/src/services/models"
	"net/http"
	"time"
)

func SendTelegramNotification(text string) {
	var resp *http.Response
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", configuration.TGConfig.BotToken)
	body, err := json.Marshal(sm.TelegramMessage{
		ParseMode: "MarkdownV2",
		Text:      text,
		ChatId:    configuration.TGConfig.ChatId,
	})
	if err != nil {
		fmt.Println(err)
	}
	for {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println(err)
		}
		if resp.StatusCode == http.StatusTooManyRequests {
			time.Sleep(time.Second * 5)
		} else {
			break
		}
	}
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(response))
	}
}

func SendUptimeNotification(projectKey, service string, state bool) {
	var status string
	if state {
		status = "active"
	} else {
		status = "error"
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02T15:04:05.999Z")
	body, err := json.Marshal(sm.UptimeNotification{
		ProjectKey: projectKey,
		Service:    service,
		Status:     status,
		Timestamp:  formattedTime,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := http.Post(
		configuration.UptimeConf.UptimeUrl,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		fmt.Println("Failed to sent uptime notification, cause: " + err.Error())
	}
	response, _ := io.ReadAll(resp.Body)
	fmt.Println(string(response))
}
