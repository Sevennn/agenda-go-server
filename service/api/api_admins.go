package api
import (
	"github.com/Sevennn/agenda-go-server/service/service"
)
func DeleteUser(uname string) bool {
	if flag := service.DeleteUser(uname); flag {
		return true
	} else {
		return false
	}
}