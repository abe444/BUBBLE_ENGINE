package model

import (
	"database/sql"

	"github.com/abe444/BUBBLE_ENGINE/util"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	row := util.Db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func CreateUser(username, password string) error {
	_, err := util.Db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	return err
}
