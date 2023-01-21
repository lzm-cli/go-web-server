package models

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/lzm-cli/gin-web-server-template/session"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateIgnoreIfExist(ctx *gin.Context, v interface{}) error {
	return session.DB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(v).Error
}

func CreateUpdateAllIfExist(ctx *gin.Context, v interface{}) error {
	return session.DB(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(v).Error
}

func RunInTransaction(ctx *gin.Context, fn func(tx *gorm.DB) error) error {
	return session.DB(ctx).Transaction(fn, &sql.TxOptions{Isolation: sql.LevelSerializable})
}
