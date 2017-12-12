package entity

import (
	"errors"
	"log"
	"agenda-go-server/service/loghelper"
)

// UserFilter : UserFilter types take an *User and return a bool value.
type UserFilter func (*User) bool
// MeetingFilter : MeetingFilter types take an *User and return a bool value.
type MeetingFilter func (*Meeting) bool

var curUserName *string;

// var dirty bool

// var uData []User
// var mData []Meeting

var errLog *log.Logger

func init()  {
	errLog = loghelper.Error
	// dirty = false
	// if err := readFromFile(); err != nil {
	// 	errLog.Println("readFromFile fail:", err)
	// }
}

// Logout : logout
func Logout() error {
	curUserName = nil
	return Sync()
}

// Sync : sync file
func Sync() error {
	// if err := writeToFile(); err != nil {
	// 	errLog.Println("writeToFile fail:", err)
	// 	return err
	// }
	return nil
}


// CreateUser : create a user
// @param a user object
func CreateUser(v *User) error {
	return insertUser(v)
	// uData = append(uData, *v.Copy())
	// dirty = true
}

// QueryUser : query users
// @param a lambda function as the filter
// @return a list of fitted users
func QueryUser(filter UserFilter) []User {
	var user []User
	uData := findAllUsers()
	for _, v := range uData {
		if filter(&v) {
			user = append(user, v)
		}
	}
	return user
}

// QueryUserByName : query user
// @param name
// @return  user
func QueryUserByName(v string) *User {
	return findUserByName(v)
}

// UpdateUser : update users
// @param a lambda function as the filter
// @param a lambda function as the method to update the user
// @return the number of updated users
func UpdateUser(filter UserFilter, switcher func (*User)) int {
	count := 0
	uData := findAllUsers()
	for i := 0; i < len(uData); i++ {
		if v := &uData[i]; filter(v) {
			origin := v.Copy()
			switcher(v)
			updateUser(origin, v)
			count++
		}
	}
	return count
}

// DeleteUser : delete users
// @param a lambda function as the filter
// @return the number of deleted users
func DeleteUser(filter UserFilter) int {
	count := 0
	uData := findAllUsers()
	length := len(uData)
	for i := 0; i < length; {
		if filter(&uData[i]) {
			length--
			deleteUser(&uData[i])
			uData[i] = uData[length]
			uData = uData[:length]
			count++
		} else {
			i++
		}
	}
	return count
}

// CreateMeeting : create a meeting
// @param a meeting object
func CreateMeeting(v *Meeting) error {
	return insertMeeting(v)
	// mData = append(mData, *v.Copy())
	// dirty = true
}

// QueryMeeting : query meetings
// @param a lambda function as the filter
// @return a list of fitted meetings
func QueryMeeting(filter MeetingFilter) []Meeting {
	var met []Meeting
	mData := findAllMeetings()
	for _, v := range mData {
		if filter(&v) {
			met = append(met, v)
		}
	}
	return met
}

// QueryMeetingByTitle : query meeting
// @param title
// @return meeting
func QueryMeetingByTitle(v string) *Meeting {
	return findMeetingByTitle(v)
}

// UpdateMeeting : update meetings
// @param a lambda function as the filter
// @param a lambda function as the method to update the meeting
// @return the number of updated meetings
func UpdateMeeting(filter MeetingFilter, switcher func (*Meeting)) int {
	count := 0
	mData := findAllMeetings()
	for i := 0; i < len(mData); i++ {
		if v := &mData[i]; filter(v) {
			origin := v.Copy()
			switcher(v)
			// loghelper.Error.Println("UpdateMeeting:", origin)
			// loghelper.Error.Println("UpdateMeeting:", v)
			updateMeeting(origin, v)
			count++
		}
	}
	return count
}

// DeleteMeeting : delete meetings
// @param a lambda function as the filter
// @return the number of deleted meetings
func DeleteMeeting(filter MeetingFilter) int {
	count := 0
	mData := findAllMeetings()
	length := len(mData)
	for i := 0; i < length; {
		if filter(&mData[i]) {
			length--
			deleteMeeting(&mData[i])
			mData[i] = mData[length]
			mData = mData[:length]
			count++
		} else {
			i++
		}
	}
	return count
}

// GetCurUser : get current user
// @return the current user
// @return error if current user does not exist
func GetCurUser() (User, error) {
	if curUserName == nil {
		return User{}, errors.New("Current user does not exist")
	}
	uData := findAllUsers()
	for _, v := range uData {
		if v.Name == *curUserName {
			return v, nil
		}
	}
	return User{}, errors.New("Current user does not exist")
}

// SetCurUser : get current user
// @param current user
func SetCurUser(u *User) {
	if u == nil {
		curUserName = nil
		return
	}
	if (curUserName == nil) {
		p := u.Name
		curUserName = &p
	} else {
		*curUserName = u.Name
	}
}

// readFromFile : read file content into memory
// @return if fail, error will be returned
func readFromFile() error {
	// uData = findAllUsers()
	// mData = findAllMeetings()
	return nil
}

// writeToFile : write file content from memory
// @return if fail, error will be returned
// func writeToFile() error {
// 	if !dirty {
// 		return nil
// 	}
// 	for i := 0; i < len(uData); i++ {
// 		deleteUser(&uData[i])
// 	}
// 	for i := 0; i < len(mData); i++ {
// 		deleteMeeting(&mData[i])
// 	}
// 	var e []error
// 	for i := 0; i < len(uData); i++ {
// 		if err := insertUser(&uData[i]); err != nil {
// 			e = append(e, err)
// 		}
// 	}
// 	for i := 0; i < len(mData); i++ {
// 		if err := insertMeeting(&mData[i]); err != nil {
// 			e = append(e, err)
// 		}
// 	}
// 	if len(e) == 0 {
// 		return nil
// 	}
// 	er := e[0]
// 	for i := 1; i < len(e); i++ {
// 		er = errors.New(er.Error() + e[i].Error())
// 	}
// 	return er
// }
