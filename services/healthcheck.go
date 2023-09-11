package services

import (
	"errors"
	"fmt"
	"monitoring-service/logger"
	"monitoring-service/models/database"
	"monitoring-service/repository"
	"monitoring-service/services/notifications"
	"net"
	"net/http"
	"net/url"
)

func HealthCheckLoop(done <-chan bool, client *http.Client, projectName string, services []database.Service) {
	// cycle allows to iterate through array infinitely
	for {
		select {
		// when channel closes, this function terminates, allows to avoid goroutine leaks
		case <-done:
			fmt.Println("Stopped checking " + projectName)
			return
		default:
			for _, v := range services {
				healthcheck(projectName, &v, client)
			}
		}
	}
}

type StatusState struct {
	StatusCode int
	Active     bool
}

func healthcheck(projectName string, service *database.Service, client *http.Client) {
	var message string
	active := repository.ServiceRepository.GetServiceState(service)
	resp, err := client.Get(service.Url)

	defer func() {
		if message != "" {
			notifications.SendNotifications(projectName, service.ServiceName, message, !active)
			repository.ServiceRepository.SetServiceState(service, !active)
		}
		if resp != nil {
			_ = resp.Body.Close()
		}
		logger.Log.Infof("%s %s checked.", projectName, service.ServiceName)
	}()

	var dnsError *net.DNSError
	if errors.As(err, &dnsError) && active {
		message = fmt.Sprintf("`%s` hostname resolution failed", projectName)
		logger.Log.Warn(err.Error())
		return
	}
	var urlError *url.Error
	if errors.As(err, &urlError) && active && urlError.Timeout() {
		message = fmt.Sprintf("🚫️WARNING🚫️\n`%s %s`\nTIMED OUT", projectName, service.ServiceName)
		return
	}

	if resp == nil {
		return
	}
	statusState := StatusState{
		StatusCode: resp.StatusCode,
		Active:     active,
	}
	switch statusState {
	case StatusState{StatusCode: http.StatusOK, Active: !active}:
		message = fmt.Sprintf("🌀GOOD NEWS🌀\n`%s %s`\nIS UP", projectName, service.ServiceName)
	case StatusState{StatusCode: http.StatusNotFound, Active: active}:
		message = fmt.Sprintf("⚠️WARNING⚠️\n`%s %s`\nIS INACCESSIBLE", projectName, service.ServiceName)
	case StatusState{StatusCode: http.StatusInternalServerError, Active: active}:
		message = fmt.Sprintf("🚫️WARNING🚫️\n`%s %s`\nRETURNED 500 STATUS CODE", projectName, service.ServiceName)
	}
}
