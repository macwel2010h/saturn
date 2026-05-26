package handlers

import (
	"html/template"
	"net/http"
	"serv-test/config"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("web/html/home.html", "web/html/t_navbar.html", "web/html/t_logo.html", "web/html/t_footer.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	flash := config.App.SessionManager.PopString(r.Context(), "flash")

	Data.Flash = flash
	err = ts.ExecuteTemplate(w, "home.html", Data)
	if err != nil {
		ServerError(w, r, err)
	}
}
