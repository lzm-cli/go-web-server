package services

import (
	"fmt"

	"github.com/fox-one/mixin-sdk-go"
	"gorm.io/gorm"
)

type Service interface {
	Run() error
}

type Hub struct {
	services map[string]Service
}

func NewHub(db *gorm.DB, client *mixin.Client) *Hub {
	hub := &Hub{services: make(map[string]Service)}

	hub.registerServices()
	return hub
}

func (hub *Hub) StartService(name string) error {
	service := hub.services[name]
	if service == nil {
		return fmt.Errorf("no service found: %s", name)
	}

	return service.Run()
}

func (hub *Hub) registerServices() {
	hub.services["scan"] = &ScanService{}
}
