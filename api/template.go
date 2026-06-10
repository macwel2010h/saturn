package handlers

import (
	"net/http"
	"serv-test/config"
	"serv-test/internal/models"
)

type templateData struct {
	User            *models.User
	Post            *models.Post
	Feed            *models.Posts
	FieldErrors     map[string]string
	UserForm        *UserForm
	SigninForm      *SigninForm
	Flash           string
	IsAuthenticated bool
}

var Data = templateData{
	User:       &models.User{},
	Post:       &models.Post{},
	Feed:       &models.Posts{},
	UserForm:   &userForm,
	SigninForm: &signinForm,
}

func PostFeedDisplay(w http.ResponseWriter, r *http.Request) {

	Data.Feed.Posts = nil

	stmt := ` SELECT * FROM posts ORDER BY created_at DESC`

	postrows, err := config.App.DB.Query(stmt)
	if err != nil {
		return
	}

	defer postrows.Close()

	for postrows.Next() {

		if err := postrows.Scan(&Data.Post.ID, &Data.Post.UserName, &Data.Post.Title, &Data.Post.Content, &Data.Post.Created_at); err != nil {
			return
		}
		Data.Feed.Posts = append(Data.Feed.Posts, *Data.Post)
	}

}
