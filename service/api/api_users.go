package api

import (
	"github.com/Sevennn/agenda-go-server/service/loghelper"
	"github.com/Sevennn/agenda-go-server/service/service"
	"github.com/Sevennn/agenda-go-server/service/entity"
)

func ListAllUser() []entity.User {
	return service.ListAllUser()
}

func UserRegister(info map[string][]string) (bool, error) {
	if info[`username`] == nil || info[`password`] == nil || info[`email`] == nil || info[`phone`] == nil {
		loghelper.Info.Println("UserRegister: Error Parameter", info)
		return false, nil
	}
	return service.UserRegister(info[`username`][0], info[`password`][0], info[`email`][0], info[`phone`][0])
}

func GetUserByName(uname string) (*entity.User) {
	return entity.QueryUserByName(uname)
}
// drop this method in homework
// func UpdateUserInfo(int id, info map[string]string) bool {
// 	res := entity.UpdateUser(
// 		func (u *entity.User) bool {
// 			return u.ID == id
// 		},
// 		func (u *entity.User) {
// 			if _,ok := info["username"];ok {
// 				u.Name = info["username"]
// 			}
// 			if _,ok := info["password"];ok {
// 				u.Password = info["password"]
// 			}
// 			if _,ok := info["email"];ok {
// 				u.Email = info["email"]
// 			}
// 			if _,ok := info["phone"];ok {
// 				u.Phone = info["phone"]
// 			}
// 		}
// 	});
// 	if res != 1 {
// 		return false
// 	} else {
// 		return true
// 	}
// }

