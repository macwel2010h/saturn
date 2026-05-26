package handlers

import (
	"html/template"
	"net/http"
	"serv-test/helpers"
)

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/createAccount.html", "web/html/t_navbar.html", "web/html/t_logo.html", "web/html/t_footer.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	Data.IsAuthenticated = helpers.IsAuthenticated(r)

	err = ts.ExecuteTemplate(w, "createAccount.html", Data)
	if err != nil {
		ServerError(w, r, err)
	}

}
