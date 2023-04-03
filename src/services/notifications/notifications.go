package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"monitoring-service/src/configuration"
	"monitoring-service/src/logger"
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
		logger.Log.Error(err)
	}
	for {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
		if err != nil {
			logger.Log.Error(err)
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
		logger.Log.Error(err)
	} else {
		logger.Log.Infow("response", string(response))
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
		logger.Log.Error(err)
		return
	}
	resp, err := http.Post(
		configuration.UptimeConf.UptimeUrl,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		logger.Log.Warnln("Failed to sent uptime notification, cause:", err.Error())
	}
	response, _ := io.ReadAll(resp.Body)
	logger.Log.Info(response)
}

func SendNotifications(projectName, serviceName, message string, active bool) {
	SendTelegramNotification(message)
	SendUptimeNotification(projectName, serviceName, !active)
}
