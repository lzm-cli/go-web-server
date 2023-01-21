package middlewares

import (
	"log"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lzm-cli/gin-web-server-template/durables"
	"github.com/lzm-cli/gin-web-server-template/handlers/user"
	"github.com/lzm-cli/gin-web-server-template/models"
	"github.com/lzm-cli/gin-web-server-template/session"
	"github.com/lzm-cli/gin-web-server-template/views"
)

var whitelist = [][2]string{
	{"GET", "^/auth$"},
	{"GET", "^/$"},
	{"GET", "^/_hc$"},
}

func CurrentUser(r *gin.Context) *models.User {
	u, _ := r.Value("u").(*models.User)
	return u
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println(2)
		header := ctx.Request.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			handleUnauthorized(ctx)
			return
		}
		log.Println(2.5)
		u, err := user.AuthenticateUserByToken(ctx, header[7:])
		if durables.CheckEmptyError(err) != nil {
			views.RenderErrorResponse(ctx, err)
		} else if u == nil {
			handleUnauthorized(ctx)
		} else {
			ctx.Set("u", u)
			log.Println()
			ctx.Next()
		}
	}
}

func handleUnauthorized(ctx *gin.Context) {
	for _, pp := range whitelist {
		if pp[0] != ctx.Request.Method {
			continue
		}
		if matched, err := regexp.MatchString(pp[1], ctx.Request.URL.Path); err == nil && matched {
			ctx.Next()
			return
		}
	}
	views.RenderErrorResponse(ctx, session.AuthorizationError())
}
