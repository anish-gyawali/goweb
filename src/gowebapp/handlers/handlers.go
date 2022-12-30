package handlers

import (
	"net/http"
)

// We need ResponseWriter for sending response, and Request for handling request
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	// If following code is emitted, the url other than '/' will be handled as same
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//passsing value to display
	h := &Home{
		HomeTitle:  "Go Lang web page",
		Name:       "Anish",
		Profession: "Developer",
		Country:    "USA",
	}
	renderTemplateHome(w, "index", h)
	if h != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Similarly, for '/task'
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	t := &Task{
		TaskNumber: "One",
		Title:      "Complete documentation",
		Due:        "jan 10, 2023",
	}
	renderTemplateTask(w, "task", t)

	if t != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
