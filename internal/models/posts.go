package models

import (
	"database/sql"
	"errors"
	"time"
)

type Post struct {
	ID         int
	UserName   string
	Title      string
	Content    string
	Created_at time.Time
}

type PostModel struct {
	DB *sql.DB
}

type Posts struct {
	Posts []Post
}

var Ps = Posts{}

func (m *PostModel) StoreCreatePost(p *Post) error {

	if m == nil || m.DB == nil {
		return errors.New("database not initialized")
	}

	stmt := ` INSERT INTO posts (username, title, content) VALUES(?,?,?)`

	_, err := m.DB.Exec(stmt, p.UserName, p.Title, p.Content)
	if err != nil {
		return err
	}

	return nil
}

func (m *PostModel) StoreGetPost(p *Post) error {
	if m == nil || m.DB == nil {
		return errors.New("database not initialized")
	}
	return nil
}
