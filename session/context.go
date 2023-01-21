package session

import (
	"github.com/fox-one/mixin-sdk-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	keyDatabase    = "1"
	keyMixinClient = "2"
)

func DB(ctx *gin.Context) *gorm.DB {
	v, _ := ctx.Get(keyDatabase)
	return v.(*gorm.DB)
}

func MixinClient(ctx *gin.Context) *mixin.Client {
	v, _ := ctx.Get(keyMixinClient)
	return v.(*mixin.Client)
}

func WithDatabase(ctx *gin.Context, database *gorm.DB) {
	database = database.WithContext(ctx)
	ctx.Set(keyDatabase, database)
}
