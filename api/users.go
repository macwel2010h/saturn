package handlers

import (
	"html/template"
	"net/http"
	"serv-test/helpers"
	"serv-test/internal/models"
	"serv-test/internal/validator"
)

type UserForm struct {
	FirstName           string `form:"firstName"`
	LastName            string `form:"lastName"`
	Username            string `form:"username"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

var userForm = UserForm{}

func CreateUser(um *models.UserModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if um == nil || um.DB == nil {
			http.Error(w, "Database config not configured", http.StatusInternalServerError)
			return
		}

		err := helpers.DecodeForm(r, &userForm)
		if err != nil {
			ClientError(w, http.StatusBadRequest)
			return
		}

		userForm.CheckField(validator.NotBlank(userForm.FirstName), "firstName", "First name can not be blank.")
		userForm.CheckField(validator.NotBlank(userForm.LastName), "lastName", "Last name can not be blank.")
		userForm.CheckField(validator.NotBlank(userForm.Username), "username", "Usename can not be blank.")
		userForm.CheckField(validator.NotBlank(userForm.Email), "email", "Email can not be blank.")
		userForm.CheckField(validator.NotBlank(userForm.Password), "password", "Password can not be blank.")
		userForm.CheckField(validator.CheckUsername(userForm.Username), "username", "Username already exist.")

		Data.FieldErrors = userForm.FieldErrors

		if userForm.Valid() {
			err := models.HashPassword(&userForm.Password)
			if err != nil {
				ServerError(w, r, err)
				return
			}

			newUser := models.User{
				FirstName: userForm.FirstName,
				LastName:  userForm.LastName,
				Username:  userForm.Username,
				Email:     userForm.Email,
				Password:  userForm.Password,
			}

			err = um.StoreCreateUser(&newUser)
			if err != nil {
				ServerError(w, r, err)
				return
			}

			http.Redirect(w, r, "/welcome", 303)

		} else {

			ts, err := template.ParseFiles("web/html/createAccount.html", "web/html/t_navbar.html", "web/html/t_logo.html")

			if err != nil {
				ServerError(w, r, err)
				return
			}
			err = ts.ExecuteTemplate(w, "createAccount.html", Data)
			if err != nil {
				ServerError(w, r, err)
			}
			userForm = UserForm{}

		}

	}

}
