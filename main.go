package main

import(
  "net/http"
  "time"
)

func main(){
  p("LiveChat", version(), "started at", config.Address)

  mux := http.NewSeverMux()
  files := http.FileSever(http.Dir(config.Static)
  mux.Handle("/static/", httpStripPrefix("/static/", files))

  mux.HandleFunc("/", index)
  mux.HandleFunc("/err", err)

  mux.HandleFunc("/login", login)
  mux.HandleFunc("/logout", logout)
  mux.HandleFunc("/signup", signup)
  mux.HandleFunc("/signup_account", signupAccount)
  mux.HandleFunc("/authenticate", authenticate)

  mux.HandleFunc("/thread/new", newThread)
  mux.HandleFunc("/thread/create", createThread)
  mux.HandleFunc("/thread/post", postThread)
  mux.HandleFunc("/thread/read", readThread)

  sever := &http.Sever{
    Addr:           config.Address,
    Handler:        mux,
    ReadTimeout:    time,Duration(config.ReadTimeout * int64(time.Second)),
    WriteTimeout:   time,Duration(config.WriteTimeout * int64(time.Second)),
    MaxHeaderBytes: 1 << 20
  }
  sever.ListenAndserve()
}

