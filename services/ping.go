package services

import (
	"fmt"
	"monitoring-service/logger"
	"monitoring-service/models/database"
	"monitoring-service/repository"
	"monitoring-service/services/notifications"
	"os/exec"
	"strings"
)

func PingLoop(done <-chan bool, projectName string, ips []database.Ip) {
	for {
		select {
		// when channel closes, this function terminates, allows to avoid goroutine leaks
		case <-done:
			fmt.Println("Stopped checking ips " + projectName)
			return
		default:
			for _, v := range ips {
				ping(projectName, &v)
			}
		}
	}
}

func ping(projectName string, ip *database.Ip) {
	pingCmd := exec.Command("ping", "-c2", ip.Ip)
	pingRes, err := pingCmd.Output()
	active := repository.IpRepository.GetIpState(ip)
	var message string

	defer func() {
		if message != "" {
			notifications.SendNotifications(projectName, "server", message, !active)
			repository.IpRepository.SetIpState(ip, !active)
		}
		logger.Log.Infof("%s %s checked", projectName, ip.Ip)
	}()

	if err != nil {
		logger.Log.Error(err)
	}
	if strings.Contains(string(pingRes), "0 received") && active {
		message = fmt.Sprintf("🚨ALERT🚨\n`%s` server is down `%s`", projectName, ip.Ip)
		return
	}
	if !strings.Contains(string(pingRes), "0 packets received") && !active {
		message = fmt.Sprintf("🌿RELIEF🌿\n`%s` server is up `%s`", projectName, ip.Ip)
		return
	}
}
