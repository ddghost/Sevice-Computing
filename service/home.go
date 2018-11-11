package service

import (
	"net/http"

	"text/template"
)

func home(w http.ResponseWriter, r *http.Request)  {

	t := template.Must(template.ParseFiles("templates/login.html"))
	t.Execute(w, nil)
}