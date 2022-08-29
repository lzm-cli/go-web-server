package main

import (
	"fmt"
	"net/http"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/<%= organization %>/<%= repo %>/config"
	"github.com/unrolled/render"
	"gorm.io/gorm"

	"github.com/dimfeld/httptreemux"
	"github.com/gorilla/handlers"
	"github.com/<%= organization %>/<%= repo %>/middlewares"
	"github.com/<%= organization %>/<%= repo %>/routes"
)

func StartHTTP(db *gorm.DB, client *mixin.Client) error {
	router := httptreemux.New()
	routes.RegisterHandlers(router)
	routes.RegisterRoutes(router)
	handler := middlewares.Authenticate(router)
	handler = middlewares.Constraint(handler)
	handler = middlewares.Context(handler, db, client, render.New())
	handler = handlers.ProxyHeaders(handler)

	return http.ListenAndServe(fmt.Sprintf(":%d", config.C.Port), handler)
}
