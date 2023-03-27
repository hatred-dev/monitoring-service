package services

import (
	"context"
	"monitoring-service/database"
	sm "monitoring-service/src/services/models"
)

// TODO THIS SHIT NEEDS REFACTORING

var SupervisorObject = &Supervisor{}

type Supervisor struct {
	channels []sm.ChannelStorage
	projects []sm.ProjectFull
}

// We have to store channels to gracefully shutdown function when reloading, this mitigates possibility of goroutine leakage

// TODO needs rewriting
func (s *Supervisor) loadProjects() {
	ctx := context.Background()
	projects, _ := database.Conn.GetProjects(ctx)
	s.projects = make([]sm.ProjectFull, len(projects))
	for i, v := range projects {
		s.projects[i].ProjectName = v.ProjectName
		ips, _ := database.Conn.GetIPsByProjectName(ctx, v.ProjectName)
		for _, ip := range ips {
			s.projects[i].Ips = append(s.projects[i].Ips, sm.Ip{
				Ip:     ip.Ip,
				Active: ip.Active.Bool,
			})
		}
		services, _ := database.Conn.GetServicesByProjectName(ctx, v.ProjectName)
		for _, service := range services {
			s.projects[i].Services = append(s.projects[i].Services, sm.Service{
				Name:   service.ServiceName,
				Url:    service.Url,
				Active: service.Active.Bool,
			})
		}
	}
}

func (s *Supervisor) loadServices() {
	s.channels = make([]sm.ChannelStorage, len(s.projects))
	for i, v := range s.projects {
		storage := sm.ChannelStorage{}
		if len(v.Ips) != 0 {
			ch := make(chan bool, 1)
			storage.PingChan = ch
			go pingLoop(ch, v.ProjectName, v.Ips)
		}
		if len(v.Services) != 0 {
			ch := make(chan bool, 1)
			storage.HealthcheckChan = ch
			go healthcheckLoop(ch, v.ProjectName, v.Services)
		}
		s.channels[i] = storage
	}
}

func (s *Supervisor) ReloadServices() {
	// shutdown every related goroutine
	for _, v := range s.channels {
		if v.PingChan != nil {
			close(v.PingChan)
		}
		if v.HealthcheckChan != nil {
			close(v.HealthcheckChan)
		}
	}
	// load projects again
	s.startServices()
}

func (s *Supervisor) startServices() {
	s.loadProjects()
	s.loadServices()
}

func StartServices() {
	SupervisorObject.startServices()
}
