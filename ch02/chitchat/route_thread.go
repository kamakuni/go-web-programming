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

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func readThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		error_message(writer, request, "Cannot read thread")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			generateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}