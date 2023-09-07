package services

import (
	"monitoring-service/models/database"
	"monitoring-service/repository"
	"net/http"
	"time"
)

var SupervisorObject = &Supervisor{}

type Supervisor struct {
	channels map[string]ChannelStorage
	client   *http.Client
}

type ChannelStorage struct {
	HealthcheckChan chan bool
	PingChan        chan bool
}

// We have to store channels to gracefully shutdown function when reloading, this mitigates possibility of goroutine leakage

func (Supervisor) getProjects() []database.Project {
	return repository.ProjectRepository.GetProjects()
}

func (s *Supervisor) SetupClient() {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DisableKeepAlives = true
	s.client = &http.Client{
		Transport: transport,
		Timeout:   time.Second * 30,
	}
}

func (s *Supervisor) CreateStorage(project database.Project) {
	storage := ChannelStorage{}
	if !project.IpsEmpty() {
		storage.StartPingJob(project.ProjectName, project.Ips)
	}
	if !project.ServicesEmpty() {
		storage.StartHealthcheckJob(s.client, project.ProjectName, project.Services)
	}

	s.channels[project.ProjectName] = storage
}

func (s *Supervisor) LoadJobs() {
	projects := s.getProjects()
	s.channels = make(map[string]ChannelStorage, len(projects))
	for _, project := range projects {
		s.CreateStorage(project)
	}
}

func (s *Supervisor) ReloadServices() {
	// shutdown every related goroutine
	for _, v := range s.channels {
		v.Close()
	}
	// load projects again
	s.LoadJobs()
}

func StartServices() {
	SupervisorObject.LoadJobs()
}

func (c *ChannelStorage) Close() {
	if c.HealthcheckChan != nil {
		close(c.HealthcheckChan)
	}
	if c.PingChan != nil {
		close(c.PingChan)
	}
}

func (c *ChannelStorage) CreatePingChannel() chan bool {
	c.PingChan = make(chan bool, 1)
	return c.PingChan
}

func (c *ChannelStorage) StartPingJob(projectName string, ips []database.Ip) {
	ch := c.CreatePingChannel()
	go PingLoop(ch, projectName, ips)
}

func (c *ChannelStorage) CreateHealthcheckChannel() chan bool {
	c.HealthcheckChan = make(chan bool, 1)
	return c.HealthcheckChan
}

func (c *ChannelStorage) StartHealthcheckJob(client *http.Client, projectName string, svc []database.Service) {
	ch := c.CreateHealthcheckChannel()
	go HealthcheckLoop(ch, client, projectName, svc)
}
