package routes

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/lzm-cli/gin-web-server-template/config"
)

func RegisterRoutes(router *gin.Engine) {
	registerUser(router)
	router.GET("/", root)
	router.GET("/_hc", healthCheck)
}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"build": config.BuildVersion + "-" + runtime.Version(),
	})
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
