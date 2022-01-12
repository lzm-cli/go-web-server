package services

import (
	"context"
	"fmt"
)

type Service interface {
	Run(context.Context) error
}

type Hub struct {
	context  context.Context
	services map[string]Service
}

func NewHub() *Hub {
	hub := &Hub{services: make(map[string]Service)}
	hub.registerServices()
	return hub
}

func (hub *Hub) StartService(name string) error {
	service := hub.services[name]
	if service == nil {
		return fmt.Errorf("no service found: %s", name)
	}

	return service.Run(hub.context)
}

func (hub *Hub) registerServices() {
	hub.services["scan"] = &ScanService{}
}
