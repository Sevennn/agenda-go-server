package api

import (
	"agenda-go-server/service/service"
	"encoding/base64"
)

func GetUserKey(username string, password string) (bool, string) {
	if (service.UserLogin(username, password)) {
		return true, base64.URLEncoding.EncodeToString([]byte(password))
	} else {
		return false, ""
	}
}

