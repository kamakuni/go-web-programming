package main

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
