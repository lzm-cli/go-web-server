package models

import (
	"database/sql"
	"time"

	"github.com/lzm-cli/gin-web-server-template/durables"
)

type Model struct {
	ID        uint         `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt sql.NullTime `json:"-"`
	UpdatedAt sql.NullTime `json:"-"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}

type FileUpload struct {
	Md5       string    `json:"md5,omitempty" gorm:"type:varchar(255);primary_key"`
	URL       string    `json:"url,omitempty" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

func (FileUpload) TableName() string {
	return "file_upload"
}

func init() {
	db := durables.NewDB()
	db.AutoMigrate(
		&User{},
		&FileUpload{},
	)
}
