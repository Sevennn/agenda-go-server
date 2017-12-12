package routes
import (
	"agenda-go-server/service/entity"
	"net/http"
	//"net/url"
	"path/filepath"
	"agenda-go-server/service/api"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"fmt"
)
func NewServer() *negroni.Negroni {
	
		formatter := render.New()
	
		n := negroni.Classic()
		mx := mux.NewRouter()
	
		initApiRoutes(mx, formatter)
	
		n.UseHandler(mx)
		return n
}

func initApiRoutes(mx *mux.Router, formatter *render.Render) {
	fmt.Println("For test: Route init")
	mx.HandleFunc("/v1/users", GetAllUserHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/users", UserRegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/users/{name:[_a-zA-Z0-9]+}", GetUserByNameHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/meetings", GetAllMeetingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/meetings", CreateMeetingHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/meetings/{title:[_a-zA-Z0-9]+}", GetMeetingByTitleHandler(formatter)).Methods("GET")
}


func GetAllUserHandler(formatter *render.Render) http.HandlerFunc {
	fmt.Println("For test")
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("For test")
		req.ParseForm()
		res := api.ListAllUser()
		formatter.JSON(w, 200, res)
	}
}

func UserRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		flag, _ := api.UserRegister(req.PostForm)
		if flag == true {
<<<<<<< HEAD
			formatter.JSON(w,201,nil) // expected a user id
			http.Redirect(w,req, "users/"+req.PostForm[`username`][0], 201)
=======
			// formatter.JSON(w,201,nil) // expected a user id
			http.Redirect(w,req, "/users/"+req.PostForm[`username`][0], http.StatusFound)

>>>>>>> sen/master
		} else {
			formatter.JSON(w,404,nil)
		}
	}
}

func GetUserByNameHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Get by name")
		req.ParseForm()
<<<<<<< HEAD
		var us *entity.User
		if name := req.Form.Get("name"); name != "" {
			us = api.GetUserByName(name)
		} else {
			vars := mux.Vars(req)
			us = api.GetUserByName(vars["name"])
		}
		if us != nil {
			r.JSON(w, 200, *us)
=======
		path := filepath.FromSlash(req.RequestURI)
		_, name := filepath.Split(path)
		fmt.Println(name)
		us := api.GetUserByName(name)
		fmt.Println(us)
		if len(us) == 1 {
			r.JSON(w, 200, us)
>>>>>>> sen/master
		} else {
			r.JSON(w,404,nil)
		}
	}
}


func GetAllMeetingHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
<<<<<<< HEAD
		if name := req.Form.Get("name"); name != "" {
			r.JSON(w, 200, api.ListAllMeetings(name))
		} else {
			r.JSON(w, 404, nil)
		}
=======
		res := api.ListAllMeetings(req.Form.Get("name"))
		r.JSON(w,200,res)
>>>>>>> sen/master
	}
}

func CreateMeetingHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		flag := api.CreateMeeting(req.PostForm)
		if flag {
<<<<<<< HEAD
			r.JSON(w, 201, nil) // expected a user id
			http.Redirect(w, req, "meetings/"+req.PostForm[`title`][0], http.StatusFound)
=======
			http.Redirect(w,req, "/meetings/"+req.PostForm[`Title`][0], http.StatusFound)
>>>>>>> sen/master
		} else {
			r.JSON(w, 404, nil)
		}
	}
}

func GetMeetingByTitleHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
<<<<<<< HEAD
		if title := req.Form.Get("title"); title != "" {
			r.JSON(w, 200, api.GetMeetingByTitle(title))
		} else {
			r.JSON(w, 404, nil)
=======
		path := filepath.FromSlash(req.RequestURI)
		_, name := filepath.Split(path)
		mt := api.GetMeetingByTitle(name)
		if len(mt) == 1 {
			r.JSON(w,200,mt[0])
		} else {
			r.JSON(w,404,nil)
>>>>>>> sen/master
		}
	}
}