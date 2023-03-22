package services

import (
	"context"
	"database/sql"
	"fmt"
	"monitoring-service/database"
	sm "monitoring-service/src/services/models"
	"monitoring-service/src/services/notifications"
	"net"
	"net/http"
	"time"
)

func healthcheck(done <-chan bool, projectName string, services []sm.Service) {
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
				var message string
				var needsNotification bool
				active := getServiceState(projectName, v.Name)
				resp, err := client.Get(v.Url)
				if active {
					if resp == nil {
						message = fmt.Sprintf("Did not receive any response from `%s`", projectName)
						needsNotification = true
					} else if err, ok := err.(net.Error); ok && err.Timeout() {
						message = fmt.Sprintf("WARNING\n`%s %s`\nTIMED OUT", projectName, v.Name)
						needsNotification = true
					} else {
						switch resp.StatusCode {
						case 500:
							message = fmt.Sprintf("WARNING\n`%s %s`\nRETURNED 500 STATUS CODE", projectName, v.Name)
						case 400:
							message = fmt.Sprintf("WARNING\n`%s %s`\nIS INACCESSIBLE", projectName, v.Name)
						}
						if resp.StatusCode != 200 {
							needsNotification = true
						}
					}
				} else {
					if resp.StatusCode == 200 {
						message = fmt.Sprintf("GOOD NEWS\n`%s %s`\nIS UP", projectName, v.Name)
						needsNotification = true
					}
				}
				if needsNotification {
					notifications.SendTelegramNotification(message)
					notifications.SendUptimeNotification(projectName, v.Name, !active)
					setServiceState(ctx, projectName, v.Name, !active)
				}
				fmt.Println(fmt.Sprintf("%s %s checked.", projectName, v.Name))
				time.Sleep(time.Millisecond * 400)
			}
		}
	}
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
	fmt.Println(err)
}

func getServiceState(projectName, serviceName string) bool {
	active, _ := database.Conn.GetServiceState(context.Background(), database.GetServiceStateParams{
		ProjectName: projectName,
		ServiceName: serviceName,
	})
	return active.Bool
}
