package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"serv-test/config"
	"serv-test/internal/models"
	"serv-test/internal/validator"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/signIn.html", "web/html/t_navbar.html", "web/html/t_logo.html")
	if err != nil {
		ServerError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "signIn.html", Data)
	if err != nil {
		ServerError(w, r, err)
	}

}

type SigninForm struct {
	validator.Validator
}

func PostSignInHandler(p *models.Post) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var um = models.UserModel{}

		if config.App.DB == nil {
			http.Error(w, "Database config not configured", http.StatusInternalServerError)
			return
		}

		err := r.ParseForm()
		if err != nil {
			fmt.Print(err)
			return
		}

		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")

		var signinForm = SigninForm{}

		signinForm.CheckField(validator.NotBlank(username), "username", "Username can not be empty.")
		signinForm.CheckField(validator.NotBlank(password), "password", "Password can not be empty.")

		if !signinForm.Valid() {
			ts, err := template.ParseFiles("web/html/signIn.html", "web/html/t_navbar.html", "web/html/t_logo.html")
			if err != nil {
				ServerError(w, r, err)
				return
			}
			err = ts.ExecuteTemplate(w, "signIn.html", Data)
			if err != nil {
				ServerError(w, r, err)
			}
		} else {

			_, err := um.CheckUserInDatabase(username, password)
			if err != nil {
				ts, err := template.ParseFiles("web/html/wrongLoginRedirect.html", "web/html/t_navbar.html", "web/html/t_logo.html")
				if err != nil {
					ServerError(w, r, err)
					return
				}
				err = ts.ExecuteTemplate(w, "wrongLoginRedirect.html", Data)
				if err != nil {
					ServerError(w, r, err)
				}
				return
			} else {

				p.UserName = username
				PostFeedDisplay(w, r)

				http.Redirect(w, r, "/home", 303)
			}
		}
	}
}
