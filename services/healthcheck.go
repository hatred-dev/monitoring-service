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
				healthCheck(projectName, &v, client)
			}
		}
	}
}

func healthCheck(projectName string, service *database.Service, client *http.Client) {
	var message string
	active := repository.ServiceRepository.GetServiceState(service)
	resp, err := requestWithRetries(client, service.Url)

	defer func() {
		if message != "" {
			notifications.SendNotifications(projectName, service.ServiceName, message, !active)
			repository.ServiceRepository.SetServiceState(service, !active)
		}
		if resp != nil {
			_ = resp.Body.Close()
		}
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

	if resp.StatusCode == http.StatusOK && !active {
		message = fmt.Sprintf("🌀GOOD NEWS🌀\n`%s %s`\nIS UP", projectName, service.ServiceName)
	}
	if resp.StatusCode == http.StatusNotFound && active {
		message = fmt.Sprintf("⚠️WARNING⚠️\n`%s %s`\nIS INACCESSIBLE", projectName, service.ServiceName)
	}
	if resp.StatusCode == http.StatusInternalServerError && active {
		message = fmt.Sprintf("🚫️WARNING🚫️\n`%s %s`\nRETURNED 500 STATUS CODE", projectName, service.ServiceName)
	}
}

func requestWithRetries(client *http.Client, url string) (*http.Response, error) {
	var resp *http.Response
	var err error
	for i := 0; i <= 3; i++ {
		resp, err = client.Get(url)
		if resp != nil && resp.StatusCode == http.StatusOK {
			break
		}
	}
	return resp, err
}
