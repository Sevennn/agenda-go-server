package routes
import (
	"net/http"
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
	mx.HandleFunc("/v1/users/{name:[_a-zA-Z]+}", GetUserByNameHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/meetings", GetAllMeetingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/meetings", CreateMeetingHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/meetings/{title:[_a-zA-Z]+}", GetMeetingByTitleHandler(formatter)).Methods("GET")
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
			formatter.JSON(w,201,nil) // expected a user id
			http.Redirect(w,nil, "users/"+req.PostForm[`username`][0], 201)
		} else {
			formatter.JSON(w,404,nil)
		}
	}
}

func GetUserByNameHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		us := api.GetUserByName(req.Form[`name`][0])
		if len(us) != 1 {
			r.JSON(w, 200, us)
		} else {
			r.JSON(w,404,nil)
		}
	}
}


func GetAllMeetingHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		res := api.ListAllMeetings(req.Form[`name`][0])
		r.JSON(w,200,res)
	}
}

func CreateMeetingHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		flag := api.CreateMeeting(req.PostForm)
		if flag {
			r.JSON(w,201,nil) // expected a user id
			http.Redirect(w,nil, "meetings/"+req.PostForm[`Title`][0], http.StatusFound)
		} else {
			r.JSON(w,404,nil)
		}
	}
}

func GetMeetingByTitleHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		r.JSON(w,200,api.GetMeetingByTitle(req.Form[`title`][0]))
	}
}