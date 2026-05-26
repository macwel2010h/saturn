package handlers

import (
	"html/template"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/welcome.html", "web/html/t_navbar.html", "web/html/t_logo.html", "web/html/t_footer.html")

	if err != nil {
		ServerError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "welcome.html", Data)
	if err != nil {
		ServerError(w, r, err)
	}

	userForm = UserForm{}
}
