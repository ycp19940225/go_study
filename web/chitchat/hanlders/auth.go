package handlers

import (
	"chichat/models"
	"fmt"
	"net/http"
)

// GET /login
// 登录页面
func Login(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth.layout", "navbar", "login")
}

// GET /signup
// 注册页面
func Signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth.layout", "navbar", "signup")
}

// POST /signup
// 注册新用户
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := models.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// 通过邮箱和密码字段对用户进行认证
// 用户认证
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := models.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if user.Password == models.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}

// GET /logout
// 用户退出
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		fmt.Println("Failed to get cookie")
		session := models.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}

	cookie2 := http.Cookie{
		Name:     "_cookie",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	}
	http.SetCookie(writer, &cookie2)

	http.Redirect(writer, request, "/", 302)
}
