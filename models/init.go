package models

import (
	"github.com/gin-gonic/gin"
	"github.com/lzm-cli/gin-web-server-template/durables"
	"github.com/lzm-cli/gin-web-server-template/session"
)

var Ctx gin.Context

func init() {
	db := durables.NewDB()
	session.WithDatabase(&Ctx, db)
	db.AutoMigrate(
		&User{},
	)
}
