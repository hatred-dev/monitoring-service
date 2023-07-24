package services

import (
	"monitoring-service/logger"
	"sync"
	"time"
)

var TimerService = NewTimerService()

type ReloadTimer struct {
	shouldReload bool
	ticker       *time.Ticker
	resetChan    chan bool
	mu           sync.Mutex
}

func NewTimerService() ReloadTimer {
	return ReloadTimer{
		shouldReload: false,
		ticker:       time.NewTicker(time.Second * 10),
		resetChan:    make(chan bool, 1),
	}
}

func (t *ReloadTimer) TriggerReset() {
	t.resetChan <- true
}

func (t *ReloadTimer) Reset() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.ticker.Reset(time.Second * 10)
	t.shouldReload = true
}

func (t *ReloadTimer) Reload() {
	if !t.shouldReload {
		return
	}
	logger.Log.Info("Reloading...")
	SupervisorObject.ReloadServices()
	t.shouldReload = false
}

func (t *ReloadTimer) Watch() {
	defer t.ticker.Stop()
	for {
		select {
		case <-t.resetChan:
			t.Reset()
		case <-t.ticker.C:
			t.Reload()
		}
	}
}
