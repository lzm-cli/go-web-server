package models

import (
	"context"
	"crypto/sha256"
	"strings"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/<%= organization %>/<%= repo %>/config"
	"github.com/<%= organization %>/<%= repo %>/session"
)

const (
	DefaultAvatar = "https://images.mixin.one/E2y0BnTopFK9qey0YI-8xV3M82kudNnTaGw0U5SU065864SsewNUo6fe9kDF1HIzVYhXqzws4lBZnLj1lPsjk-0=s128"
)

func AuthenticateUserByOAuth(ctx context.Context, authorizationCode string) (string, error) {
	accessToken, scope, err := mixin.AuthorizeToken(ctx, config.C.Mixin.ClientID, config.C.Mixin.ClientSecret, authorizationCode, "")
	if err != nil {
		if strings.Contains(err.Error(), "Forbidden") {
			return "", session.ForbiddenError(ctx)
		}
		return "", session.BadDataError(ctx)
	}
	if !strings.Contains(scope, "PROFILE:READ") {
		return "", session.ForbiddenError(ctx)
	}
	me, err := mixin.UserMe(ctx, accessToken)
	if err != nil {
		return "", err
	}
	if me == nil {
		return "", session.BadDataError(ctx)
	}
	user, err := checkAndWriteUser(ctx, me.UserID, accessToken, me.FullName, me.AvatarURL, me.IdentityNumber, me.Biography)
	if err != nil {
		return "", err
	}
	authenticationToken, err := generateAuthenticationToken(user.UserId, user.AccessToken)
	if err != nil {
		return "", session.BadDataError(ctx)
	}
	return authenticationToken, nil
}

func generateAuthenticationToken(userId string, accessToken string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        userId,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	sum := sha256.Sum256([]byte(accessToken))
	return token.SignedString(sum[:])
}

func checkAndWriteUser(ctx context.Context, userId, accessToken, fullName, avatarURL, identityNumber, biography string) (*User, error) {
	if _, err := uuid.FromString(userId); err != nil {
		return nil, session.BadDataError(ctx)
	}
	if avatarURL == "" {
		avatarURL = DefaultAvatar
	}
	user := &User{
		UserId:         userId,
		FullName:       fullName,
		IdentityNumber: identityNumber,
		AvatarURL:      avatarURL,
		AccessToken:    accessToken,
	}
	if err := createUpdateAllIfExist(ctx, user); err != nil {
		return nil, session.TransactionError(ctx, err)
	}
	return user, nil
}
