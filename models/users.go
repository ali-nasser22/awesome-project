package models

import (
	"awesomeProject/db"
	"awesomeProject/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func (u User) ValidateCredentials() error {

	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievesPassword string
	err := row.Scan(&retrievesPassword)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(retrievesPassword), []byte(u.Password))
	if err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}
