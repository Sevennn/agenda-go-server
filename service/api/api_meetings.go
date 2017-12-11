package api

import (
	"agenda-go-server/service/service"
	"agenda-go-server/service/entity"
)

func ListAllMeetings(uname string) ([]entity.Meeting) {
	return entity.QueryMeeting(func (m *entity.Meeting) bool {
		return m.Sponsor == uname || m.IsParticipator(uname)
	})
}

func CreateMeeting(info map[string][]string) (bool) {
	return service.CreateMeeting(info[`Sponsor`][0], info[`title`][0], info[`StartDate`][0], info[`EndDate`][0], info[`Participators`])
}

func GetMeetingByTitle(title string) []entity.Meeting {
	return entity.QueryMeeting(func (m *entity.Meeting) bool {
		return m.Title == title
	})

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


// func DeleteMeeting(mid int) bool {
// 	return entity.DeleteMeeting(func (m *entity.Meeting) bool {
// 		return m.ID == mid
// 	}) == 1;
// }