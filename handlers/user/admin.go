package user

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/gin-gonic/gin"
	"github.com/lzm-cli/gin-web-server-template/durables"
	"github.com/lzm-cli/gin-web-server-template/models"
	"github.com/lzm-cli/gin-web-server-template/session"
	"github.com/lzm-cli/gin-web-server-template/tools"
	"gorm.io/gorm"
)

func UploadFile(c *gin.Context, u *models.User) (string, error) {
	_file, _ := c.FormFile("file")
	file, _ := _file.Open()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	key := fmt.Sprintf("%x", md5.Sum(data))

	var f models.FileUpload
	err = session.DB(c).Take(&f, "md5=?", key).Error
	if err == nil {
		return f.URL, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	a, err := durables.GetMixinClient().CreateAttachment(context.Background())
	if err != nil {
		return "", err
	}
	if err := mixin.UploadAttachment(context.Background(), a, data); err != nil {
		return "", err
	}
	if err := session.DB(c).Create(&models.FileUpload{
		Md5: key,
		URL: a.ViewURL,
	}).Error; err != nil {
		tools.Log(err)
	}
	return a.ViewURL, nil
}
