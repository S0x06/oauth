package request

import (
	"encoding/json"
	"fmt"
	//	"html/template"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/go-session/session"
	"github.com/oauth/v2/service/model"
	"github.com/oauth/v2/service/util"
	//	"gopkg.in/gin-gonic/gin.v1"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {

	store, err := session.Start(nil, w, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if req.Method == "POST" {

		req.ParseForm()
		username := req.FormValue("username")
		password := req.FormValue("password")

		user, err := model.GetUser(username, password)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user.ID > 0 {
			store.Set("uid", user.ID)
			store.Save()

			outputHTML(w, req, "../static/auth.html")
			return

		}

	}

	outputHTML(w, req, "../static/login.html")

	//	t, _ := template.ParseFiles("../static/login.html")
	//	t.Execute(w, nil)

	//	outputHTML(w, r, "../static/login.html")
}

func AuthHandler(w http.ResponseWriter, req *http.Request) {

	store, err := session.Start(nil, w, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uid, ok := store.Get("uid")
	if !ok {
		outputHTML(w, req, "../static/login.html")
		return
	}
	fmt.Println("type:", reflect.TypeOf(uid))
	//	var strUid string
	//	strUid = strconv.Itoa(uid)
	//	strUid, _, err = util.RedisGet(strUid)
	//	//	//	auth := a.(AuthRequest)

	//	fmt.Println(strUid)

}

type AuthRequest struct {
	ClientId     string `json:"client_id"`
	ResponseType string `json:"response_type"`
	RedirectUri  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
	State        string `json:"state"`
}

func AuthorizeHandler(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	client_id := req.FormValue("client_id")
	response_type := req.FormValue("response_type")
	redirect_uri := req.FormValue("redirect_uri")
	scope := req.FormValue("scope")
	state := req.FormValue("state")

	if client_id == "" {
		http.Error(w, "授权失败", http.StatusInternalServerError)
		return
	}

	if response_type != "code" {
		http.Error(w, "授权失败", http.StatusInternalServerError)
		return
	}

	if redirect_uri == "" {
		http.Error(w, "授权失败", http.StatusInternalServerError)
		return
	}

	if scope == "" {
		http.Error(w, "授权失败", http.StatusInternalServerError)
		return
	}

	if state == "" {
		http.Error(w, "授权失败", http.StatusInternalServerError)
		return
	}

	client, err := model.GetAuthorize(client_id)
	if err != nil {
		http.Error(w, "授权失败", http.StatusInternalServerError)
		return
	}

	if client.RedirectUri != redirect_uri {
		http.Error(w, "该服务器未授权", http.StatusInternalServerError)
		return
	}

	store, err := session.Start(nil, w, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uid, ok := store.Get("uid")
	if !ok {
		outputHTML(w, req, "../static/login.html")
		//		w.Header().Set("Location", "/login")
		//		w.WriteHeader(http.StatusFound)
		return
	}

	auth := new(AuthRequest)

	auth.ClientId = client_id
	auth.ResponseType = response_type
	auth.RedirectUri = redirect_uri
	auth.Scope = scope
	auth.State = state

	authJson, err := json.Marshal(auth)
	if err != nil {
		fmt.Println("error:", err)
	}

	//	time_out := time.Now().ParseDuration("10m")
	strUid := strconv.Itoa(uid)
	util.RedisSave(strUid, authJson, 10*time.Second)

	//	w.Header().Set("Location", "/auth")
	//	w.WriteHeader(http.StatusFound)

	outputHTML(w, req, "../static/auth.html")
	return
}

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
