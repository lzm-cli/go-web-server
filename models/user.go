package models

import (
	"time"
)

type User struct {
	UserId         string    `json:"user_id" gorm:"primary_key;type:varchar(36);"`
	IdentityNumber string    `json:"identity_number" gorm:"type:varchar(11);"`
	FullName       string    `json:"full_name" gorm:"type:varchar;"`
	AvatarURL      string    `json:"avatar_url" gorm:"type:varchar;"`
	AccessToken    string    `json:"access_token" gorm:"type:varchar;"`
	CreatedAt      time.Time `json:"created_at" gorm:"default: now()"`
}

func (User) TableName() string {
	return "users"
}
