package middlewares

import (
	"log"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/gin-gonic/gin"
	"github.com/lzm-cli/gin-web-server-template/session"
	"gorm.io/gorm"
)

func Context(db *gorm.DB, client *mixin.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println(3)
		session.WithDatabase(ctx, db)
		ctx.Next()
	}
}
