package user

import (
	"crypto/sha256"
	"strings"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/lzm-cli/gin-web-server-template/config"
	"github.com/lzm-cli/gin-web-server-template/models"
	"github.com/lzm-cli/gin-web-server-template/session"
)

func AuthenticateUserByOAuth(ctx *gin.Context, authorizationCode string) (string, error) {
	accessToken, scope, err := mixin.AuthorizeToken(ctx, config.C.Mixin.ClientID, config.C.Mixin.ClientSecret, authorizationCode, "")
	if err != nil {
		if strings.Contains(err.Error(), "Forbidden") {
			return "", session.ForbiddenError()
		}
		return "", session.BadDataError()
	}
	if !strings.Contains(scope, "PROFILE:READ") {
		return "", session.ForbiddenError()
	}
	me, err := mixin.UserMe(ctx, accessToken)
	if err != nil {
		return "", err
	}
	if me == nil {
		return "", session.BadDataError()
	}
	user, err := checkAndWriteUser(ctx, me.UserID, accessToken, me.FullName, me.AvatarURL, me.IdentityNumber, me.Biography)
	if err != nil {
		return "", err
	}
	authenticationToken, err := generateAuthenticationToken(user.UserId)
	if err != nil {
		return "", session.BadDataError()
	}
	return authenticationToken, nil
}

func generateAuthenticationToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        userId,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	sum := sha256.Sum256([]byte(config.C.Key))
	return token.SignedString(sum[:])
}

func checkAndWriteUser(ctx *gin.Context, userId, accessToken, fullName, avatarURL, identityNumber, biography string) (*models.User, error) {
	if _, err := uuid.FromString(userId); err != nil {
		return nil, session.BadDataError()
	}
	user := &models.User{
		UserId:         userId,
		FullName:       fullName,
		IdentityNumber: identityNumber,
		AvatarURL:      avatarURL,
		AccessToken:    accessToken,
	}
	if err := models.CreateUpdateAllIfExist(ctx, user); err != nil {
		return nil, session.TransactionError(err)
	}
	return user, nil
}
