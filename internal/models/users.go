package models

import (
	"database/sql"
	"errors"
	"serv-test/config"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	CreatedAt []uint8
}

type UserModel struct {
	DB *sql.DB
}

func (um *UserModel) StoreCreateUser(u *User) error {

	stmt := `INSERT INTO users (firstName, lastName, username, email, password)
	VALUES(?,?,?,?,?)`

	_, err := um.DB.Exec(stmt, u.FirstName, u.LastName, u.Username, u.Email, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (um *UserModel) StoreDeleteUser(username string) (int, error) {
	if um == nil || um.DB == nil {
		return 0, errors.New("database not initialized")
	}
	return 0, nil
}

func (um *UserModel) CheckUserInDatabase(username, password string) (User, error) {

	var u = User{}
	stmt := `SELECT * FROM users WHERE username = ?`

	err := um.DB.QueryRow(stmt, username).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, errors.New("No user found")
		}
		return User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return User{}, errors.New("wrong password.")
	}
	return u, nil
}
func (um *UserModel) CheckUsernameAvailability(usernameForm string) bool {
	if um == nil || um.DB == nil {
		return false
	}
	stmt := `SELECT username FROM users`

	usernamerows, err := um.DB.Query(stmt)
	if err != nil {
		return false
	}

	defer usernamerows.Close()

	var matched bool = false

	for usernamerows.Next() {

		var username string

		usernamerows.Scan(&username)
		if username == usernameForm {
			matched = true
		}
	}
	return matched
}

func CheckUsernameAvailability(usernameForm string) bool {
	um := UserModel{DB: config.App.DB}
	return um.CheckUsernameAvailability(usernameForm)
}

func HashPassword(pass *string) error {
	if pass == nil {
		return errors.New("The password is nil.")
	}

	passBytes := []byte(*pass)

	hashedBytes, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	*pass = string(hashedBytes[:])
	return nil
}
