package models

import (
	"time"
)

type User struct {
	Model
	UserId         string    `json:"user_id" gorm:"type:varchar(36);index:,unique;"`
	IdentityNumber string    `json:"identity_number" gorm:"type:varchar(11);index:,unique;"`
	FullName       string    `json:"full_name" gorm:"type:varchar;"`
	AvatarURL      string    `json:"avatar_url" gorm:"type:varchar;"`
	AccessToken    string    `json:"access_token" gorm:"type:varchar;"`
	CreatedAt      time.Time `json:"created_at" gorm:"type:timestamp;default: now()"`
}

func (User) TableName() string {
	return "users"
}
