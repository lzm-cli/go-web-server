package models

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/<%= organization %>/<%= repo %>/session"
	"github.com/<%= organization %>/<%= repo %>/tools"
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

func AuthenticateUserByToken(ctx context.Context, authenticationToken string) (*User, error) {
	var ua string
	var user *User
	var queryErr error
	token, err := jwt.Parse(authenticationToken, func(token *jwt.Token) (interface{}, error) {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, session.BadDataError(ctx)
		}

		_, ok = token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, session.BadDataError(ctx)
		}
		user, queryErr = FindUserById(ctx, fmt.Sprint(claims["jti"]))
		if queryErr != nil {
			return nil, queryErr
		}
		if user == nil {
			return nil, session.BadDataError(ctx)
		}
		sum := sha256.Sum256([]byte(user.UserId + ua))
		return sum[:], nil
	})

	if queryErr != nil {
		return nil, queryErr
	}
	if err != nil || !token.Valid {
		return nil, nil
	}
	return user, nil
}

func FindUserById(ctx context.Context, userId string) (*User, error) {
	var user User
	err := db.First(&user, "user_id=?", userId).Error
	if CheckEmptyError(err) != nil {
		return nil, err
	}
	if user.UserId == "" {
		u, err := session.MixinClient(ctx).ReadUser(ctx, userId)
		if err != nil {
			return nil, err
		}
		user = User{
			UserId:         u.UserID,
			IdentityNumber: u.IdentityNumber,
			FullName:       u.FullName,
			AvatarURL:      u.AvatarURL,
			CreatedAt:      u.CreatedAt,
		}
		if err := db.Create(&user).Error; err != nil {
			tools.Log(err)
		}
		return nil, nil
	}
	return &user, nil
}
