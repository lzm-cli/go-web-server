package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Constraint() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(1)
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,Mixin-Conversation-ID")
		c.Header("Access-Control-Allow-Methods", "OPTIONS,GET,PUT,POST,DELETE")
		c.Header("Access-Control-Max-Age", "600")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
