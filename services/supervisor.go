package services

import (
	"monitoring-service/models/database"
	"monitoring-service/repository"
	"net/http"
	"os"
	"sync"
	"time"
)

type ChannelStorage []chan bool

var supervisorInstance *Supervisor
var supervisorOnce sync.Once

func GetSupervisor() *Supervisor {
	supervisorOnce.Do(func() {
		supervisorInstance = &Supervisor{
			mappedStorage:   make(map[string]ChannelStorage),
			channelsStorage: make(ChannelStorage, 0),
		}
		supervisorInstance.SetupClient()
	})
	return supervisorInstance
}

type Supervisor struct {
	channelsStorage ChannelStorage
	mappedStorage   map[string]ChannelStorage
	client          *http.Client
}

func (*Supervisor) getProjects() []database.Project {
	return repository.ProjectRepository.GetProjects()
}

func (s *Supervisor) SetupClient() {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DisableKeepAlives = true
	timeout := time.Second * 30 // Default timeout value
	if customTimeout, ok := os.LookupEnv("CUSTOM_TIMEOUT"); ok {
		if parsedTimeout, err := time.ParseDuration(customTimeout); err == nil {
			timeout = parsedTimeout
		}
	}
	s.client = &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
}

func (s *Supervisor) LoadProjectJobs(project database.Project) {
	s.mappedStorage[project.ProjectName] = make(ChannelStorage, 0)
	if !project.IpsEmpty() {
		s.StartPingJob(project.ProjectName, project.Ips)
	}
	if !project.ServicesEmpty() {
		s.StartHealthCheckJob(project.ProjectName, project.Services)
	}
}

func (s *Supervisor) StartProjectJobs() {
	projects := s.getProjects()
	for _, project := range projects {
		s.LoadProjectJobs(project)
	}
}

func (s *Supervisor) ReloadProjectJobs() {
	// shutdown every related goroutine
	for _, v := range s.mappedStorage {
		for _, ch := range v {
			close(ch)
		}
	}
	// load jobs again
	s.StartProjectJobs()
}

func StartServices() {
	GetSupervisor().StartProjectJobs()
}

func (s *Supervisor) AllocateChannel() chan bool {
	ch := make(chan bool, 1)
	s.channelsStorage = append(s.channelsStorage, ch)
	return ch
}

func (s *Supervisor) AllocateMappedChannel(key string) chan bool {
	if _, exists := s.mappedStorage[key]; !exists {
		s.mappedStorage[key] = make(ChannelStorage, 0)
	}
	ch := make(chan bool, 1)
	s.mappedStorage[key] = append(s.mappedStorage[key], ch)
	return ch
}

func (s *Supervisor) StartPingJob(projectName string, ips []database.Ip) {
	go PingLoop(s.AllocateMappedChannel(projectName), projectName, ips)
}

func (s *Supervisor) StartHealthCheckJob(projectName string, svc []database.Service) {
	go HealthCheckLoop(s.AllocateMappedChannel(projectName), s.client, projectName, svc)
}
