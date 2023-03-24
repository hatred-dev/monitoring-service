package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"monitoring-service/database"
	sm "monitoring-service/src/services/models"
	"monitoring-service/src/services/notifications"
	"net"
	"net/http"
	"net/url"
	"time"
)

func healthcheckLoop(done <-chan bool, projectName string, services []sm.Service) {
	ctx := context.Background()
	client := &http.Client{
		Timeout: time.Second * 15,
	}
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
			time.Sleep(time.Second * 1)
		}
	}
}

func healthcheck(projectName string, service *sm.Service, client *http.Client, ctx context.Context) {
	var dnsError *net.DNSError
	var message string
	active := getServiceState(ctx, projectName, service.Name)
	resp, err := client.Get(service.Url)

	defer func() {
		if message != "" {
			sendNotifications(projectName, service.Name, message, !active)
			setServiceState(ctx, projectName, service.Name, !active)
		}
		if resp != nil {
			err := resp.Body.Close()
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println(fmt.Sprintf("%s %s checked.", projectName, service.Name))
	}()

	if errors.As(err, &dnsError) && active {
		message = fmt.Sprintf("`%s` hostname resolution failed", projectName)
		return
	}
	if err, ok := err.(*url.Error); ok && active && err.Timeout() {
		message = fmt.Sprintf("WARNING\n`%s %s`\nTIMED OUT", projectName, service.Name)
		return
	}
	if resp == nil {
		return
	}
	if resp.StatusCode == 500 && active {
		message = fmt.Sprintf("🚫️WARNING🚫️\n`%s %s`\nRETURNED 500 STATUS CODE", projectName, service.Name)
		return
	}
	if resp.StatusCode == 404 && active {
		message = fmt.Sprintf("⚠️WARNING⚠️\n`%s %s`\nIS INACCESSIBLE", projectName, service.Name)
		return
	}
	if resp.StatusCode == 200 && !active {
		message = fmt.Sprintf("🌀GOOD NEWS🌀\n`%s %s`\nIS UP", projectName, service.Name)
		return
	}

}

func sendNotifications(projectName, serviceName, message string, active bool) {
	notifications.SendTelegramNotification(message)
	notifications.SendUptimeNotification(projectName, serviceName, !active)
}

func setServiceState(ctx context.Context, projectName, serviceName string, active bool) {
	err := database.Conn.SetServiceState(ctx, database.SetServiceStateParams{
		ProjectName: projectName,
		ServiceName: serviceName,
		Active: sql.NullBool{
			Bool:  active,
			Valid: true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}

func getServiceState(ctx context.Context, projectName, serviceName string) bool {
	active, _ := database.Conn.GetServiceState(ctx, database.GetServiceStateParams{
		ProjectName: projectName,
		ServiceName: serviceName,
	})
	return active.Bool
}
