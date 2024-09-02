package services

import (
    "monitoring-service/logger"
    "sync"
    "time"
)

var TimerService = NewTimerService(GetSupervisor())

type ReloadTimer struct {
    shouldReload bool
    ticker       *time.Ticker
    resetChan    chan bool
    service      ServiceInterface
    mu           sync.Mutex
}

type ServiceInterface interface {
    ReloadProjectJobs()
}

func NewTimerService(service ServiceInterface) *ReloadTimer {
    return &ReloadTimer{
        shouldReload: false,
        ticker:       time.NewTicker(time.Second * 10),
        resetChan:    make(chan bool, 1),
        service:      service,
    }
}

func (t *ReloadTimer) Reload() {
    t.mu.Lock()
    defer t.mu.Unlock()
    if !t.shouldReload {
        return
    }
    logger.Log.Info("Reloading...")
    t.service.ReloadProjectJobs()
    t.shouldReload = false
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
