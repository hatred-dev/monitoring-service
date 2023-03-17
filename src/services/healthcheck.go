package services

import (
	"context"
	"database/sql"
	"fmt"
	"monitoring-service/database"
	"net/http"
	"net/url"
	"time"
)

func healthcheck(done <-chan bool, projectName string, services []Service) {
	ctx := context.Background()
	client := http.Client{
		Timeout: time.Second * 15,
	}
	// cycle allows to iterate through array over and over
	for {
		select {
		case <-done:
			fmt.Sprintln("Stopped checking", projectName)
			return
		default:
			for _, v := range services {
				active := getServiceState(projectName, v.Name)
				resp, err := client.Get(v.Url)
				if err != nil {
					urlError := err.(*url.Error)
					if urlError.Timeout() && active {
						text := fmt.Sprintf("WARNING\n`%s %s`\nTIMED OUT", projectName, v.Name)
						setServiceState(ctx, projectName, v.Name, false)
						sendNotification(text)
					}
				}
				if resp.StatusCode == 200 && !active {
					text := fmt.Sprintf("GOOD NEWS\n`%s %s`\nIS UP", projectName, v.Name)
					setServiceState(ctx, projectName, v.Name, true)
					sendNotification(text)
				}
				if resp.StatusCode == 500 && active {
					text := fmt.Sprintf("WARNING\n`%s %s`\nRETURNED 500 STATUS CODE", projectName, v.Name)
					setServiceState(ctx, projectName, v.Name, false)
					sendNotification(text)
				}
				if resp.StatusCode == 400 && active {
					text := fmt.Sprintf("WARNING\n`%s %s`\nIS INACCESSIBLE", projectName, v.Name)
					setServiceState(ctx, projectName, v.Name, false)
					sendNotification(text)
				}
				fmt.Println(fmt.Sprintf("%s %s checked.", projectName, v.Name))
				time.Sleep(time.Millisecond * 500)
			}
		}
	}
}

func setServiceState(ctx context.Context, projectName, serviceName string, active bool) {
	database.Conn.SetServiceState(ctx, database.SetServiceStateParams{
		ProjectName: projectName,
		ServiceName: serviceName,
		Active: sql.NullBool{
			Bool:  active,
			Valid: true,
		},
	})
}

func getServiceState(projectName, serviceName string) bool {
	active, _ := database.Conn.GetServiceState(context.Background(), database.GetServiceStateParams{
		ProjectName: projectName,
		ServiceName: serviceName,
	})
	return active.Bool
}
