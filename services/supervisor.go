package services

import (
	"monitoring-service/models/api"
	"monitoring-service/models/database"
	"monitoring-service/repository"
	"net/http"
	"time"
)

// TODO THIS SHIT NEEDS REFACTORING

var SupervisorObject = &Supervisor{}

type Supervisor struct {
	channels map[string]api.ChannelStorage
	client   *http.Client
}

// We have to store channels to gracefully shutdown function when reloading, this mitigates possibility of goroutine leakage

func loadProjects() []database.Project {
	return repository.ProjectRepository.GetProjects()
}

func (s *Supervisor) startMonitoring() {
	projects := loadProjects()
	s.channels = make(map[string]api.ChannelStorage, len(projects))
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DisableKeepAlives = true
	s.client = &http.Client{
		Transport: transport,
		Timeout:   time.Second * 30,
	}
	for _, project := range projects {
		storage := api.ChannelStorage{}
		if len(project.Ips) != 0 {
			ch := make(chan bool, 1)
			storage.PingChan = ch
			go pingLoop(ch, project.ProjectName, project.Ips)
		}
		if len(project.Services) != 0 {
			ch := make(chan bool, 1)
			storage.HealthcheckChan = ch
			go healthcheckLoop(ch, s.client, project.ProjectName, project.Services)
		}
		s.channels[project.ProjectName] = storage
	}
}

func (s *Supervisor) ReloadServices() {
	// shutdown every related goroutine
	for _, v := range s.channels {
		if v.PingChan != nil {
			v.PingChan <- true
		}
		if v.HealthcheckChan != nil {
			v.HealthcheckChan <- true
		}
	}
	// load projects again
	s.startMonitoring()
}

func StartServices() {
	SupervisorObject.startMonitoring()
}
