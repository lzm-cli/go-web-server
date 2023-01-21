package main

import (
	"fmt"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/lzm-cli/gin-web-server-template/config"
	"github.com/lzm-cli/gin-web-server-template/middlewares"
	"github.com/lzm-cli/gin-web-server-template/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartHTTP(db *gorm.DB, client *mixin.Client) error {
	// router := httptreemux.New()
	// routes.RegisterHandlers(router)
	// routes.RegisterRoutes(router)
	// handler := middlewares.Authenticate(router)
	// handler = middlewares.Constraint(handler)
	// handler = middlewares.Context(handler, db, client, render.New())
	// handler = handlers.ProxyHeaders(handler)
	// return http.ListenAndServe(fmt.Sprintf(":%d", config.C.Port), handler)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(middlewares.Constraint())
	r.Use(middlewares.Authenticate())
	r.Use(middlewares.Context(db, client))
	// routes.RegisterHandlers(r)
	routes.RegisterRoutes(r)
	r.Run(fmt.Sprintf(":%d", config.C.Port))
	return nil
}
