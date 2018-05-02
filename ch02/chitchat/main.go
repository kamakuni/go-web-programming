package main

import (
	"net/http"
	"text/template"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	/*
		mux.HandleFunc("/", index)
		mux.HandleFunc("/err", err)
		mux.HandleFunc("/login", login)
		mux.HandleFunc("/logout", logout)
		mux.HandleFunc("/signup", signup)
		mux.HandleFunc("/signup_account", aignupAccount)
		mux.HandleFunc("authenticate", authenticate)

		mux.HandleFunc("/thread/new", newThread)
		mux.HandleFunc("/thread/create", createThread)
		mux.HandleFunc("/thread/post", postThread)
		mux.HandleFunc("/thread/add", addThread)*/
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
