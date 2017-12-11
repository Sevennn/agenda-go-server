package api

import (
	"agenda-go-server/service"
)

func GetUserKey(username string, password string) (bool, string) {
	if (service.UserLogin(username, password)) {
		return true, base64.URLEncoding.EncodeToString(password)
	} else {
		return false, ""
	}
}

