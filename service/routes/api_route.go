package routes
import (
	//"github.com/Sevennn/agenda-go-server/service/entity"
	"net/http"
	//"net/url"
	"path/filepath"
	"github.com/Sevennn/agenda-go-server/service/api"
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
			formatter.JSON(w,201,nil) // expected a user id
			http.Redirect(w,req, "users/"+req.PostForm[`username`][0], 201)
		} else {
			formatter.JSON(w,404,nil)
		}
	}
}

func GetUserByNameHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Get by name")
		req.ParseForm()
		path := filepath.FromSlash(req.RequestURI)
		_, name := filepath.Split(path)
		fmt.Println(name)
		us := api.GetUserByName(name)
		fmt.Println(us)
		if us != nil {
			r.JSON(w, 200, us)
		} else {
			r.JSON(w,404,nil)
		}
	}
}


func GetAllMeetingHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if name := req.Form.Get("name"); name != "" {
			r.JSON(w, 200, api.ListAllMeetings(name))
		} else {
			r.JSON(w, 404, nil)
		}
	}
}

func CreateMeetingHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		flag := api.CreateMeeting(req.PostForm)
		if flag {
			r.JSON(w, 201, nil) // expected a user id
			http.Redirect(w, req, "meetings/"+req.PostForm[`Title`][0], http.StatusFound)
		} else {
			r.JSON(w, 404, nil)
		}
	}
}

func GetMeetingByTitleHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		path := filepath.FromSlash(req.RequestURI)
		_, name := filepath.Split(path)
		mt := api.GetMeetingByTitle(name)
		if len(mt) == 1 {
			r.JSON(w,200,mt[0])
		} else {
			r.JSON(w,404,nil)
		}
	}
}