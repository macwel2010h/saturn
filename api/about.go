package handlers

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/about.html", "web/html/t_navbar.html", "web/html/t_logo.html", "web/html/t_footer.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "about.html", Data)
	if err != nil {
		ServerError(w, r, err)
	}

}
