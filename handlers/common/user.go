package common

import "github.com/lzm-cli/gin-web-server-template/config"

func CheckIsAdmin(userID string) bool {
	return config.C.Admin[userID]
}
