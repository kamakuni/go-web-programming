package main

import (
	"net/http"
)

// GET /threads/new
// Show the new thread from page
func newThread(writer http.ResponseWriter, request *http.Requet) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// Create the user account
func createThread(writer http.ResponseWriter, request *http.Requet) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} esle {
		err = request.ParseForm()
		if err != nil {
			// danger
		}
		topic := request.PostFromValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			// danger
		}
		http.Redirect(writer, request, "/", 302)
	}
}