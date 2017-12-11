package api

import (
	"agenda-go-server/service"
)

func ListAllMeetings(int uid) (bool, []entity.Meeting) {
	u := entity.QueryUser(func (u *entity.User) bool {
		return u.ID == uid
	})
	if len(u) == 0 {
		return false,nil
	}
	return true, entity.QueryMeeting(func (m *entity.Meeting) {
		return m.Sponsor == u.Name || m.IsParticipator(u.Name)
	})
}

func CreateMeeting(info map[string][]string) (bool,int) {
	u := entity.QueryUser(func (u *entity.User) bool {
		return u.ID == info[`uid`]
	})
	sponsor := u.Name
	participator := info[`Participators`]
	title := info[`Title`]
	startdate := info[`StartDate`]
	enddate := info[`EndDate`]
	for _, i := range participator {
		if username == i {
			errLog.Println("Create Meeting: sponsor can't be participator")
			return false,0
		}
		l := entity.QueryUser(func (u *entity.User) bool{
			return u.Name == i
		})
		if (len(l) == 0) {
			errLog.Println("Create Meeting: no such a user : ", i)
			return false,0
		}
		dc := 0
		for _, j := range participator {
			if j == i {
				dc++
				if dc == 2 {
					errLog.Println("Create Meeting: duplicate participator")
					return false,0
				}
			}
		}
	}
	sTime,err := entity.StringToDate(startDate)
	if err != nil {
		errLog.Println("Create Meeting: Wrong Date")
		return false,0
	}
	eTime,err := entity.StringToDate(endDate)
	if err != nil {
		errLog.Println("Create Meeting: Wrong Date")
		return false,0
	}
	if eTime.LessThan(sTime) == true {
		errLog.Println("Create Meeting: Start Time greater than end time")
		return false,0
	}
	for _, p := range participator {
		l := entity.QueryMeeting(func (m *entity.Meeting) bool {
			if m.Sponsor == p || m.IsParticipator(p) {
				if m.StartDate.LessOrEqual(sTime) && m.EndDate.MoreThan(sTime) {
					return true
				}
				if m.StartDate.LessThan(eTime) && m.EndDate.GreateOrEqual(eTime) {
					return true
				}
				if m.StartDate.GreateOrEqual(sTime) && m.EndDate.LessOrEqual(eTime) {
					return true
				}
			}
			return false
		})
		if len(l) > 0 {
			errLog.Println("Create Meeting: ",p," time conflict")
			return false,0
		}
	}
	tu := entity.QueryUser(func (u *entity.User) bool {
		return u.Name == username
	})
	if len(tu) == 0 {
		errLog.Println("Create Meeting: Sponsor ", username, " not exist")
		return false,0
	}
	l := entity.QueryMeeting(func (m *entity.Meeting) bool {
		if m.Sponsor == username || m.IsParticipator(username) {
			if m.StartDate.LessOrEqual(sTime) && m.EndDate.MoreThan(sTime) {
				return true
			}
			if m.StartDate.LessThan(eTime) && m.EndDate.GreateOrEqual(eTime) {
				return true
			}
			if m.StartDate.GreateOrEqual(sTime) && m.EndDate.LessOrEqual(eTime) {
				return true
			}
		}
		return false
	})

	if len(l) > 0 {
		errLog.Println("Create Meeting: ", username, " time conflict")
		return false,0
	}
	mid := entity.CreateMeeting(&entity.Meeting{username, participator,sTime,eTime, title})
	return true,mid
}

func GetMeetingByID(mid int) entity.Meeting {
	m := entity.QueryMeeting(func (m *entity.Meeting) {
		return m.ID == mid
	})
	if len(m) != 1 {
		return {}
	} else {
		return m[0]
	}
}

// func UpdateMeeting(mid int, uid int, info map[string][]string) (bool, int) {
// 	uname := info[`Sponsor`]
// 	u := entity.QueryUser(func (u *entity.User) bool {
// 		return u.ID == uid && u.Name == uname
// 	})
// 	if len(u) != 1 {
// 		return false, 0
// 	}

// 	if StartDate,ok := info["StartDate"];ok {
// 		if sTime,err := entity.StringToDate(StartDate); err != nil {
// 			errLog.Println("Update Meeting: Wrong Date")
// 			return false,0
// 		}
// 	}

// 	if EndDate,ok := info["EndDate"];ok {
// 		if eTime,err := entity.StringToDate(EndDate); err != nil {
// 			errLog.Println("Update Meeting: Wrong Date")
// 			return false,0
// 		}
// 	}


// 	count := entity.UpdateMeeting(
// 		func (m *entity.Meeting) bool {
// 			return m.ID == mid
// 		},
// 		func (m *entity.Meeting) {
// 			if _, ok := info["Participators"];ok {
// 				m.Participators = info["Participators"]
// 			}
// 			if _, ok := info["Title"];ok {
// 				m.Title = info["Title"]
// 			}
// 			if _, ok := info["StartDate"];ok {
// 				m.StartDate = info["StartDate"]
// 			}
// 			if _, ok := info["EndDate"];ok {
// 				m.EndDate = info["EndDate"]
// 			}
// 		}
// 	)
// }


func DeleteMeeting(mid int) bool {
	return entity.DeleteMeeting(func (m *entity.Meeting) bool {
		return m.ID == mid
	}) == 1;
}