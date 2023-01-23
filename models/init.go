package models

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lzm-cli/gin-web-server-template/durables"
	"github.com/lzm-cli/gin-web-server-template/session"
)

var Ctx gin.Context

type Model struct {
	ID        uint         `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt sql.NullTime `json:"-"`
	UpdatedAt sql.NullTime `json:"-"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}

type FileUpload struct {
	Md5       string    `json:"key,omitempty" gorm:"type:varchar(255);primary_key"`
	URL       string    `json:"url,omitempty" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

func (FileUpload) TableName() string {
	return "file_upload"
}

func init() {
	db := durables.NewDB()
	session.WithDatabase(&Ctx, db)
	db.AutoMigrate(
		&User{},
	)
}
