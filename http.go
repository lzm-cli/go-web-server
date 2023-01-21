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
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(middlewares.Constraint())
	r.Use(middlewares.Authenticate())
	r.Use(middlewares.Context(db, client))
	routes.RegisterRoutes(r)
	r.Run(fmt.Sprintf(":%d", config.C.Port))
	return nil
}
