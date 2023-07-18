package services

import (
	"context"
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

func healthcheckLoop(done <-chan bool, client *http.Client, projectName string, services []database.Service) {
	ctx := context.Background()
	// cycle allows to iterate through array infinitely
	for {
		select {
		// when channel closes, this function terminates, allows to avoid goroutine leaks
		case <-done:
			fmt.Println("Stopped checking " + projectName)
			return
		default:
			for _, v := range services {
				healthcheck(projectName, &v, client, ctx)
			}
		}
	}
}

func healthcheck(projectName string, service *database.Service, client *http.Client, ctx context.Context) {
	var dnsError *net.DNSError
	var message string
	active := repository.ProjectRepository.GetServiceState(projectName, service.ServiceName)
	resp, err := client.Get(service.Url)

	defer func() {
		if message != "" {
			notifications.SendNotifications(projectName, service.ServiceName, message, !active)
			repository.ProjectRepository.SetServiceState(projectName, service.ServiceName, !active)
		}
		if resp != nil {
			err := resp.Body.Close()
			if err != nil {
				fmt.Println(err)
			}
		}
		logger.Log.Infof("%s %s checked.", projectName, service.ServiceName)
	}()

	if errors.As(err, &dnsError) && active {
		message = fmt.Sprintf("`%s` hostname resolution failed", projectName)
		logger.Log.Warn(err.Error())
		return
	}
	if err, ok := err.(*url.Error); ok && active && err.Timeout() {
		message = fmt.Sprintf("🚫️WARNING🚫️\n`%s %s`\nTIMED OUT", projectName, service.ServiceName)
		return
	}
	if resp == nil {
		return
	}
	if resp.StatusCode == 500 && active {
		message = fmt.Sprintf("🚫️WARNING🚫️\n`%s %s`\nRETURNED 500 STATUS CODE", projectName, service.ServiceName)
		return
	}
	if resp.StatusCode == 404 && active {
		message = fmt.Sprintf("⚠️WARNING⚠️\n`%s %s`\nIS INACCESSIBLE", projectName, service.ServiceName)
		return
	}
	if resp.StatusCode == 200 && !active {
		message = fmt.Sprintf("🌀GOOD NEWS🌀\n`%s %s`\nIS UP", projectName, service.ServiceName)
		return
	}
}
