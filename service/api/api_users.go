package api

import (
	"service"
)

func ListAllUser() []entity.User {
	return entity.QueryUser(func (u *entity.User) bool {
		return true
	})
}

func UserRegister(info map[string][]string) (bool, int) {
	username := info[`username`]
	password := info[`password`]
	email := info[`email`]
	phone := info[`phone`]
	user := entity.QueryUser(func (u *entity.User) bool {
		return u.Name == username
	})
	if len(user) == 1 {
		errLog.Println("User Register: Already exist username")
		return false, 0
	}
	uid := entity.CreateUser(&entity.User{username, password, email, phone})

	return true, uid
}

func GetUserByID(int id) (entity.User) {
	res := entity.QueryUser(func (u *entity.User) bool {
		return u.ID == id
	})
	if len(res) != 1 {
		return nil
	} else {
		return res[0]
	}
}

func UpdateUserInfo(int id, info map[string]string) bool {
	res := entity.UpdateUser(
		func (u *entity.User) bool {
			return u.ID == id
		},
		func (u *entity.User) {
			if _,ok := info["username"];ok {
				u.Name = info["username"]
			}
			if _,ok := info["password"];ok {
				u.Password = info["password"]
			}
			if _,ok := info["email"];ok {
				u.Email = info["email"]
			}
			if _,ok := info["phone"];ok {
				u.Phone = info["phone"]
			}
		}
	});
	if res != 1 {
		return false
	} else {
		return true
	}
}

