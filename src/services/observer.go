package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"monitoring-service/database"
	"monitoring-service/src/configuration"
	"net/http"
)

// TODO THIS SHIT NEEDS REFACTORING

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

func sendNotification(text string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", configuration.TGConfig.BotToken)
	body, err := json.Marshal(TelegramMessage{
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

// We have to store channels to gracefully shutdown function when reloading, this mitigates possibility of goroutine leakage
var channels map[string]ChannelStorage

// TODO needs rewriting
func loadProjects() []ProjectFull {
	ctx := context.Background()
	projects, _ := database.Conn.GetProjects(ctx)
	projectsArr := make([]ProjectFull, len(projects))
	for i, v := range projects {
		projectsArr[i].ProjectName = v.ProjectName
		ips, _ := database.Conn.GetIPsByProjectName(ctx, v.ProjectName)
		for _, ip := range ips {
			projectsArr[i].Ips = append(projectsArr[i].Ips, Ip{
				Ip:     ip.Ip,
				Active: ip.Active.Bool,
			})
		}
		services, _ := database.Conn.GetServicesByProjectName(ctx, v.ProjectName)
		for _, service := range services {
			projectsArr[i].Services = append(projectsArr[i].Services, Service{
				Name:   service.ServiceName,
				Url:    service.Url,
				Active: service.Active.Bool,
			})
		}
	}
	return projectsArr
}

func executeServices(projects []ProjectFull) {
	for _, v := range projects {
		storage := ChannelStorage{}
		if len(v.Ips) != 0 {
			storage.PingChan = nil
		}
		if len(v.Services) != 0 {
			ch := make(chan bool, 1)
			storage.HealthcheckChan = ch
			go healthcheck(ch, v.ProjectName, v.Services)
		}
		channels[v.ProjectName] = storage
	}
}

func ReloadServices() {
	for _, v := range channels {
		if v.PingChan != nil {
			close(v.PingChan)
		}
		if v.HealthcheckChan != nil {
			close(v.HealthcheckChan)
		}
	}
	executeServices(loadProjects())
}

func StartServices() {
	channels = make(map[string]ChannelStorage)
	executeServices(loadProjects())
}
