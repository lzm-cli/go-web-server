package services

import (
	"context"
	"fmt"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/<%= organization %>/<%= repo %>/session"
	"gorm.io/gorm"
)

type Service interface {
	Run(context.Context) error
}

type Hub struct {
	context  context.Context
	services map[string]Service
}

func NewHub(db *gorm.DB, client *mixin.Client) *Hub {
	hub := &Hub{services: make(map[string]Service)}
	hub.context = session.WithDatabase(context.Background(), db)
	hub.context = session.WithMixinClient(hub.context, client)

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
