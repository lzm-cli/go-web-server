package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lzm-cli/gin-web-server-template/handlers/user"
	"github.com/lzm-cli/gin-web-server-template/middlewares"
	"github.com/lzm-cli/gin-web-server-template/session"
	"github.com/lzm-cli/gin-web-server-template/views"
)

func registerAdminUser(r *gin.RouterGroup) {
	b := &adminUsersImpl{}
	r.POST("/upload", b.uploadFile)
}

type adminUsersImpl struct{}

func (b *adminUsersImpl) uploadFile(c *gin.Context) {
	if url, err := user.UploadFile(c, middlewares.CurrentUser(c)); err != nil {
		views.RenderErrorResponse(c, session.BadDataError())
	} else {
		views.RenderDataResponse(c, map[string]string{"view_url": url})
	}
}
