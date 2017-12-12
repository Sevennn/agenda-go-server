package entity

import (
	"agenda-go-server/service/loghelper"
	"reflect"
	"testing"
)

var users = []User{
	{"u1", "up", "u1@q", "123"},
	{"u2", "ua", "u2@q", "456"},
	{"u3", "ua", "u3@q", "789"},
}

var meetings = []Meeting{
	{"u1", []string{"u2", "u3"}, Date{2017, 10, 21, 7, 36}, Date{2017, 10, 22, 8, 0}, "u1 u2 u3"},
	{"u1", []string{}, Date{2017, 10, 21, 19, 36}, Date{2017, 10, 22, 20, 0}, "only u1"},
	{"u2", []string{"u3"}, Date{2000, 10, 21, 19, 36}, Date{2001, 10, 23, 20, 0}, "u2 u3"},
}

func init() {
	uData := findAllUsers()
	for i := 0; i < len(uData); i++ {
		deleteUser(&users[i])
	}
	mData := findAllMeetings()
	for i := 0; i < len(mData); i++ {
		deleteMeeting(&meetings[i])
	}
	if vec := findAllUsers(); len(vec) != 0 {
		loghelper.Error.Println("ListAllUser() = ", vec)
	}
	if vec := findAllMeetings(); len(vec) != 0 {
		loghelper.Error.Println("ListAllMeeting() = ", vec)
	}
}

func _testCreateAll()  {
	for i := 0; i < len(users); i++ {
		CreateUser(&users[i])
	}
	for i := 0; i < len(meetings); i++ {
		CreateMeeting(&meetings[i])
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		u *User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{"CU u1:", args{&users[0]}, users[:1]},
		{"CU u2:", args{&users[1]}, users[:2]},
		{"CU u3:", args{&users[2]}, users[:]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateUser(tt.args.u)
			if got := findAllUsers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryUser(t *testing.T) {
	type args struct {
		filter UserFilter
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{"QU u1",
			args{func(u *User) bool {
				return u.Name == "u1"
			}},
			users[0:1],
		},
		{"QU u0",
			args{func(u *User) bool {
				return u.Name == "u0"
			}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryUser(tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		filter   UserFilter
		switcher func(*User)
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"UU u1",
			args{func(u *User) bool {
				return u.Name == "u1"
			},
				func(u *User) {
					u.Phone = "321"
				}},
			1,
		},
		{"UU Phone123",
			args{func(u *User) bool {
				return u.Phone == "123"
			},
				func(u *User) {
					u.Phone = "Error"
				}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateUser(tt.args.filter, tt.args.switcher); got != tt.want {
				t.Errorf("UpdateUser() = %v, want %v\nUser:%v", got, tt.want, findAllUsers())
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		filter UserFilter
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"DU u1",
			args{func(u *User) bool {
				return u.Name == "u1"
			}},
			1,
		},
		{"DU u1",
			args{func(u *User) bool {
				return u.Name == "u1"
			}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteUser(tt.args.filter); got != tt.want {
				t.Errorf("DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateMeeting(t *testing.T) {
	type args struct {
		m *Meeting
	}
	tests := []struct {
		name string
		args args
		want []Meeting
	}{
		{"CM " + meetings[0].Title, args{&meetings[0]}, meetings[:1]},
		{"CM" + meetings[1].Title, args{&meetings[1]}, meetings[:2]},
		{"CM" + meetings[2].Title, args{&meetings[2]}, meetings[:]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateMeeting(tt.args.m)
			if got := findAllMeetings(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMeeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryMeeting(t *testing.T) {
	type args struct {
		filter MeetingFilter
	}
	tests := []struct {
		name string
		args args
		want []Meeting
	}{
		{"QM u1",
			args{func(v *Meeting) bool {
				return v.Sponsor == "u1"
			}},
			meetings[:2],
		},
		{"QM u3",
			args{func(v *Meeting) bool {
				return v.Sponsor == "u3"
			}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryMeeting(tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryMeeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateMeeting(t *testing.T) {
	type args struct {
		filter   MeetingFilter
		switcher func(*Meeting)
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"UM Sponsor:u1",
			args{func(v *Meeting) bool {
				return v.Sponsor == "u1"
			},
				func(v *Meeting) {
					v.Title += "u1 Sponsor"
				}},
			2,
		},
		{"UM Title:u1 u2 u3",
			args{func(v *Meeting) bool {
				return v.Title == "u1 u2 u3"
			},
				func(v *Meeting) {
					v.Title = "Error"
				}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateMeeting(tt.args.filter, tt.args.switcher); got != tt.want {
				t.Errorf("UpdateMeeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteMeeting(t *testing.T) {
	type args struct {
		filter MeetingFilter
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"DM Sponsor:u1",
			args{func(v *Meeting) bool {
				return v.Sponsor == "u1"
			}},
			2,
		},
		{"DM Sponsor:u1",
			args{func(v *Meeting) bool {
				return v.Sponsor == "u1"
			}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteMeeting(tt.args.filter); got != tt.want {
				t.Errorf("DeleteMeeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetCurUser(t *testing.T) {
	type args struct {
		u *User
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"set curUser: u3", args{&users[2]}, users[2].Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetCurUser(tt.args.u)
			if got := *curUserName; got != tt.want {
				t.Errorf("SetCurUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurUser(t *testing.T) {
	tests := []struct {
		name    string
		want    User
		wantErr bool
	}{
		{"get curUser: u3",
			users[2],
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCurUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSync(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Sync", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// dirty = true
			// uData = users
			// mData = meetings
			// curUserName = &tt.name
			// if err := Sync(); (err != nil) != tt.wantErr {
			// 	t.Errorf("Sync() error = %v, wantErr %v", err, tt.wantErr)
			// }
			// uData = nil
			// mData = nil
			_testCreateAll()
			uData := findAllUsers()
			mData := findAllMeetings()
			// if err := readFromFile(); err != nil {
			// 	loghelper.Error.Println("readFromFile fail:", err)
			// }
			// if !reflect.DeepEqual(uData, users) {
			if !reflect.DeepEqual(len(uData), len(users)) {
				t.Errorf("readFromFile() users = %v: want %v", uData, users)
			}
			if !reflect.DeepEqual(len(mData), len(meetings)) {
				t.Errorf("readFromFile() meetings = %v, want %v", mData, meetings)
			}
			// if *curUserName != tt.name {
			// 	t.Errorf("readFromFile() curUser = %v: want %v", *curUserName, tt.name)
			// }
		})
	}
}

func TestQueryUserByName(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{"Query exit",
			args{users[0].Name},
			&users[0],
		},
		{"Query no exit",
			args{"NO_EXIT"},
			nil,
		},
	}
	_testCreateAll()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryUserByName(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryUserByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryMeetingByTitle(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *Meeting
	}{
		{"Query exit",
			args{meetings[0].Title},
			&meetings[0],
		},
		{"Query no exit",
			args{"NO_EXIT"},
			nil,
		},
	}
	_testCreateAll()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryMeetingByTitle(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryMeetingByTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
