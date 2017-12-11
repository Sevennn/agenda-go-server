package service
import (
	"net/http"
	"agenda-go-server/service/api"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"strconv"
)

func initApiRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/v1/users", GetAllUserHandler(formatter)).Method("GET")
	mx.HandleFunc("/v1/users", UserRegisterHandler(formatter)).Method("POST")
	mx.HandleFunc("/v1/users/{id:[0-9]+}", GetUserByIDHandler(formatter)).Method("GET")
	mx.HandleFunc("/v1/meetings", GetAllMeetingHandler(formatter)).Method("GET")
	mx.HandleFunc("/v1/meetings", CreateHandler(formatter)).Method("POST")
	mx.HandleFunc("/v1/meetings/{id:[0-9]+}", GetMeetingByIDHandler(formatter)).Method("GET")
}


func GetAllUserHandler(formatter *render.Render) http.HandlerFunc {	
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		res := api.ListAllUser()
		formatter.JSON(w, 200, res)
	}
}

func UserRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		flag, uid := api.UserRegister(req.PostForm)
		if flag {
			formatter.JSON(w,201,{}) // expected a user id
			http.Redirect(w,nil, "users/"+strconv.Itoa(uid), 201)
		} else {
			formatter.JSON(w,404,{})
		}
	}
}

func GetUserByIDHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		u := api.GetUserByID(req.Form[`id`])
		if u != nil {
			r.JSON(w, 200, u)
		} else {
			r.JSON(w,404,{})
		}
	}
}


func GetAllMeetingHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		res := api.ListAllMeetings(req.Form[`id`])
		r.JSON(w,200,res)
	}
}

func CreateMeetingHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		flag,mid := api.CreateMeeting(req.PostForm)
		if flag {
			formatter.JSON(w,201,{}) // expected a user id
			http.Redirect(w,nil, "meetings/"+strconv.Itoa(mid), http.StatusFound)
		} else {
			formatter.JSON(w,404,{})
		}
	}
}

func GetMeetingByIDHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		r.JSON(w,200,api.GetMeetingByID(req.Form[`mid`]))
	}
}