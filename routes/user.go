package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lzm-cli/gin-web-server-template/handlers/user"
	"github.com/lzm-cli/gin-web-server-template/middlewares"
	"github.com/lzm-cli/gin-web-server-template/views"
)

func registerUser(router *gin.RouterGroup) {
	impl := &usersImpl{}

	router.GET("/auth", impl.authenticate)
	router.GET("/me", impl.me)
}

type usersImpl struct{}

func (impl *usersImpl) authenticate(c *gin.Context) {
	code := c.Query("code")
	if token, err := user.AuthenticateUserByOAuth(c, code); err != nil {
		views.RenderErrorResponse(c, err)
	} else {
		views.RenderDataResponse(c, token)
	}
}

func (impl *usersImpl) me(c *gin.Context) {
	views.RenderDataResponse(c, middlewares.CurrentUser(c))
}
