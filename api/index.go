package handlers

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/index.html", "web/html/t_navbar.html", "web/html/t_logo.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "index.html", Data)
	if err != nil {
		ServerError(w, r, err)
	}

}
