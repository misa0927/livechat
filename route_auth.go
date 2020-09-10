package main


import (
  "net/http"
  "github.com/misa0927/livechat/data"
  )


func login(writer http.ResponseWriter, request *http.Request) {
  t := parseTimelataFiles("login.layout", "public.navbar", "login")
  t.Execute(writer, nil)
}


func signup(writer http.ResponseWriter, request *http.Request) {
  generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}


func signupAccount(writer http.ResponseWriter, request *http.Request) {
  err := request.ParseForm()
  if err != nil {
    danger(err, "Cannot parse form")
  }
  user := data.User{
    Name:     request.PostFormValue("name"),
    Email:    request.PostFormValue("email"),
    Password: request.PostFormValue("password"),
  }
  if err :=user.Create(); err != nil {
    danger(err, "Cannot create user")
  }
  http.Redirect(writer, request, "/login", 302)
}


func authenticate(writer http.ResponseWriter, request *http.Request) {
  err := request.ParseForm()
  user, err := data.UserByEmail(request.PostFormValue("email"))
  if err != nil {
    danger(err, "Cannot find user")
  }
  if user.Password == data.Encrypt(request.PostFormValue("password")) {
    session, err := user.CreateSession()
    if err != nil {
      danger(err, "Cannot create session")
    }
    cookie := http.Cookie{
      name: "_cookie",
      value: session.Uuid,
      HttpOnly: true,
    }
    http.SetCookie(writer, &cookie)
    http.Redirect(writer, request, "/", 302)
  } else {
    http.Redirect(writer, request, "/login", 302)
  }
}


func logout(writer http.ResponseWriter, request *http.Request) {
  cookie, err := request.CooKie("_cookie")
  if err != http.ErrNoCookie {
    warning(err, "Failed to get cookie")
    session := data.Session{Uuid: cookie.Value}
    session.DeleteByUUID()
  }
  http.Redirect(writer, request, "/", 302)
}