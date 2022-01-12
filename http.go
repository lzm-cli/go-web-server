package main

import (
	"fmt"
	"net/http"

	"github.com/<%= organization %>/<%= repo %>/config"
	"github.com/unrolled/render"

	"github.com/<%= organization %>/<%= repo %>/middlewares"
	"github.com/<%= organization %>/<%= repo %>/routes"
	"github.com/dimfeld/httptreemux"
	"github.com/gorilla/handlers"
)

func StartHTTP() error {
	router := httptreemux.New()
	routes.RegisterHanders(router)
	routes.RegisterRoutes(router)
	handler := middlewares.Authenticate(router)
	handler = middlewares.Constraint(handler)
	handler = middlewares.Context(handler, render.New())
	handler = handlers.ProxyHeaders(handler)

	return http.ListenAndServe(fmt.Sprintf(":%d", config.C.Port), handler)
}
