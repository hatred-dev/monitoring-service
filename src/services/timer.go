package services

import (
	"monitoring-service/src/logger"
	"time"
)

var TimerService = NewTimerService()

type ReloadTimer struct {
	shouldReload bool
	ticker       *time.Ticker
	resetChan    chan bool
}

func NewTimerService() ReloadTimer {
	return ReloadTimer{
		shouldReload: false,
		ticker:       time.NewTicker(time.Second * 10),
		resetChan:    make(chan bool, 1),
	}
}

func (t *ReloadTimer) ResetTimer() {
	t.resetChan <- true
}

func (t *ReloadTimer) StartTimer() {
	for {
		select {
		case <-t.resetChan:
			t.ticker.Reset(time.Second * 10)
			t.shouldReload = true
		case <-t.ticker.C:
			if t.shouldReload {
				logger.Log.Info("Reloading...")
				SupervisorObject.ReloadServices()
				t.shouldReload = false
			}
		}
	}
}
