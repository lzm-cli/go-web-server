package services

import (
	"fmt"

	"github.com/lzm-cli/gin-web-server-template/durables"
	"github.com/lzm-cli/gin-web-server-template/models"
	"github.com/lzm-cli/gin-web-server-template/tools"
)

type ScanService struct{}

func (service *ScanService) Run() error {
	fmt.Println("test")
	var u models.User
	durables.GetDB().First(&u)
	tools.PrintJson(u)
	return nil
}
