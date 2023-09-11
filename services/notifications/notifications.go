package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"monitoring-service/configuration"
	"monitoring-service/logger"
	"monitoring-service/models/api"
	"net/http"
	"time"
)

func SendTelegramNotification(text string) error {
	var resp *http.Response
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", configuration.TGConfig.BotToken)
	body, err := json.Marshal(api.TelegramMessage{
		ParseMode: "MarkdownV2",
		Text:      text,
		ChatId:    configuration.TGConfig.ChatId,
	})
	if err != nil {
		return err
	}
	for {
		resp, err = http.Post(url, "application/json", bytes.NewBuffer(body))
		if err != nil {
			return err
		}
		if resp.StatusCode == http.StatusTooManyRequests {
			time.Sleep(time.Second * 5)
		} else {
			break
		}
	}
	var response []byte
	if resp != nil {
		response, err = io.ReadAll(resp.Body)
	}
	if err != nil {
		return err
	}
	logger.Log.Info(string(response))
	return nil
}

func SendUptimeNotification(projectKey, service string, state bool) error {
	status := "error"
	if state {
		status = "active"
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02T15:04:05.999Z")
	body, err := json.Marshal(api.UptimeNotification{
		ProjectKey: projectKey,
		Service:    service,
		Status:     status,
		Timestamp:  formattedTime,
	})
	if err != nil {
		return err
	}
	resp, err := http.Post(
		configuration.UptimeConf.UptimeUrl,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}
	response, _ := io.ReadAll(resp.Body)
	logger.Log.Info(response)
	return nil
}

func SendNotifications(projectName, serviceName, message string, active bool) {
	err := SendTelegramNotification(message)
	if err != nil {
		logger.Log.Warn(err)
	}
	err = SendUptimeNotification(projectName, serviceName, active)
	if err != nil {
		logger.Log.Warn(err)
	}
}
