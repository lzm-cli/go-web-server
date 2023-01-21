package views

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lzm-cli/gin-web-server-template/session"
	"github.com/lzm-cli/gin-web-server-template/tools"
	"gorm.io/gorm"
)

type ResponseView struct {
	Data  interface{} `json:"data,omitempty"`
	Error error       `json:"error,omitempty"`
	Prev  string      `json:"prev,omitempty"`
	Next  string      `json:"next,omitempty"`
}

func RenderDataResponse(c *gin.Context, view interface{}) {
	c.JSON(http.StatusOK, ResponseView{Data: view})
	c.Abort()
}

func RenderErrorResponse(c *gin.Context, err error) {
	sessionError, ok := err.(session.Error)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		sessionError = session.ValidationError("record not found")
	} else if !ok {
		sessionError = session.ServerError(err)
		tools.Log(err)
		tools.PrintJson(c.Request)
	}
	if sessionError.Code == 10001 {
		sessionError.Code = 500
	}
	c.JSON(sessionError.Status, ResponseView{Error: sessionError})
	c.Abort()
}

func RenderBlankResponse(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseView{})
	c.Abort()
}
