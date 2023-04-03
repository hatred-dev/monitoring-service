package services

import (
	"context"
	"database/sql"
	"fmt"
	"monitoring-service/database"
	"monitoring-service/src/logger"
	sm "monitoring-service/src/services/models"
	"monitoring-service/src/services/notifications"
	"os/exec"
	"strings"
	"time"
)

func pingLoop(done <-chan bool, projectName string, ips []sm.Ip) {
	ctx := context.Background()
	for {
		select {
		// when channel closes, this function terminates, allows to avoid goroutine leaks
		case <-done:
			fmt.Println("Stopped checking ips " + projectName)
			return
		default:
			for _, v := range ips {
				ping(ctx, projectName, &v)
			}
			time.Sleep(time.Second * 1)
		}
	}
}

func ping(ctx context.Context, projectName string, ip *sm.Ip) {
	pingCmd := exec.Command("ping", ip.Ip)
	pingRes, err := pingCmd.Output()
	active, _ := database.Conn.GetIpState(ctx, ip.Ip)
	var message string

	defer func() {
		if message != "" {
			notifications.SendNotifications(projectName, "server", message, !active.Bool)
			setIpState(ctx, ip, !active.Bool)
		}
		logger.Log.Infof("%s %s checked", projectName, ip.Ip)
	}()

	if err != nil {
		logger.Log.Error(err)
	}
	if strings.Contains(string(pingRes), "0 packets received") && active.Bool {
		message = fmt.Sprintf("🚨ALERT🚨\n`%s` server is down `%s`", projectName, ip.Ip)
		return
	}
	if !strings.Contains(string(pingRes), "0 packets received") && !active.Bool {
		message = fmt.Sprintf("🌿RELIEF🌿\n`%s` server is up `%s`", projectName, ip.Ip)
		return
	}
}

func setIpState(ctx context.Context, ip *sm.Ip, active bool) {
	err := database.Conn.SetIpState(ctx, database.SetIpStateParams{
		Active: sql.NullBool{
			Bool:  active,
			Valid: true,
		},
		Ip: ip.Ip,
	})
	if err != nil {
		logger.Log.Error(err)
	}
}
