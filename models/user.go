package models

import (
	"errors"

	"github.com/zin-min-thu/go-rest-api/db"
	"github.com/zin-min-thu/go-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hasshePassowrd, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hasshePassowrd)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&retrievedPassword)
	if err != nil {
		return err
	}

	// compare password
	passwordIsValid := utils.CheckPassword(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
