package api

func DeleteUser(uid int) {
	return entity.DeleteUser(func (u *entity.User) bool {
		return u.ID == uid 
	}) == 1;
}