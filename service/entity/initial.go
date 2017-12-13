package entity

import (
	// _ : for init
	_ "github.com/mattn/go-sqlite3"
	// _ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
	"github.com/go-xorm/xorm"
	"github.com/Sevennn/agenda-go-server/service/loghelper"
	"path/filepath"
)

var orm *xorm.Engine
// var logFile *os.File

var dbFilePath = "/src/github.com/Sevennn/agenda-go-server/service/data/db.db"

func init() {
	dbFilePath = filepath.Join(loghelper.GoPath, dbFilePath)

    var err error
	if orm, err = xorm.NewEngine("sqlite3", dbFilePath); err != nil {
		loghelper.Error.Println("Fail to create engine: ", err)
	}
	// orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	// orm.SetMapper(core.SameMapper{})

	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// orm.SetDefaultCacher(cacher)
	// if logFile, err = os.Create("mysql.log"); err != nil {
	// 	panic(err)
	// }
	// orm.SetLogger(xorm.NewSimpleLogger(logFile))
	// defer orm.Close()
	// orm.ShowSQL(true)

	if err := orm.Sync(new(User)); err != nil {
		loghelper.Error.Println("Fail to sync database: ", err)
	}
	if err := orm.Sync(new(Met)); err != nil {
		loghelper.Error.Println("Fail to sync database: ", err)
	}
}

func insertUser(v *User) error {
	if affected, err := orm.Insert(v); err != nil {
		loghelper.Error.Println("insertUser Error:", affected, err)
		return err
	}
	return nil
}

func deleteUser(v *User) error {
	if affected, err := orm.Delete(v); err != nil {
		loghelper.Error.Println("deleteUser Error:", affected, err)
		return err
	}
	return nil
}

func updateUser(origin, modify *User) error {
	if affected, err := orm.Update(modify, origin); err != nil {
		loghelper.Error.Println("updateUser Error:", affected, err)
		return err
	}
	return nil
}

func findAllUsers() []User {
	vec := make([]User, 0)
	if err := orm.Find(&vec); err != nil {
		loghelper.Error.Println("findAllUsers Error:", err)
	}
	return vec
}

func findUserByName(name string) *User {
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	u := &User{Name: name}
	has, err := orm.Get(u)
	if err != nil {
		loghelper.Error.Println("findUserByName Error:", err)
	}
	if has {
		return u
	}
	return nil
}

func insertMeeting(vv *Meeting) error {
	v := vv.toMet()
	// loghelper.Error.Println("insertMeeting = ", vv, v)
	if affected, err := orm.Insert(v); err != nil {
		loghelper.Error.Println("insertMeeting Error:", affected, err)
		return err
	}
	return nil
}

func findAllMeetings() []Meeting {
	vec := make([]Met, 0)
	if err := orm.Find(&vec); err != nil {
		loghelper.Error.Println("findAllMeetings Error:", err)
	}
	out := make([]Meeting, len(vec))
	for i := 0; i < len(vec); i++ {
		out[i] = *vec[i].toMet()
		// loghelper.Error.Println("findAllMeetings = ", vec[i], out[i])
	}
	return out
}

func findMeetingByTitle(title string) *Meeting {
	v := &Met{Title: title}
	has, err := orm.Get(v)
	if err != nil {
		loghelper.Error.Println("findMeetingByTitle Error:", err)
	}
	if has {
		return v.toMet()
	}
	return nil
}

func updateMeeting(origin, modify *Meeting) error {
	if affected, err := orm.Delete(origin.toMet()); err != nil {
		loghelper.Error.Println("deleteMeeting Error:", affected, err)
		return err
	}
	if affected, err := orm.Insert(modify.toMet()); err != nil {
		loghelper.Error.Println("insertMeeting Error:", affected, err)
		return err
	}
	// if affected, err := orm.Update(m, o); err != nil {
	// 	loghelper.Error.Println("updateMeeting Error:", affected, err)
	// 	return err
	// } else {
	// 	loghelper.Error.Println("updateMeeting affected:", affected)
	// }
	return nil
}

func deleteMeeting(vv *Meeting) error {
	v := vv.toMet()
	if affected, err := orm.Delete(v); err != nil {
		loghelper.Error.Println("deleteMeeting Error:", affected, err)
		return err
	}
	return nil
}
