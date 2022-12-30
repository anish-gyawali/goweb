package handlers

import (
	"html/template"
	"net/http"
)

//parse template files

var homeTemplates = template.Must(template.ParseFiles("templates/index.html"))
var taskTemplates = template.Must(template.ParseFiles("templates/task.html"))

// declaration of struct need for the Home template
type Home struct {
	HomeTitle  string
	Name       string
	Profession string
	Country    string
}

// declaration of struct need for the Task template
type Task struct {
	TaskNumber string
	Title      string
	Due        string
}

// A custom render function which takes the file name of template html file
func renderTemplateHome(w http.ResponseWriter, tmp string, h *Home) {
	err := homeTemplates.ExecuteTemplate(w, tmp+".html", h)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func renderTemplateTask(w http.ResponseWriter, tmp string, t *Task) {
	err := taskTemplates.ExecuteTemplate(w, tmp+".html", t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
