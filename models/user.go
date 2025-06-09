package models

import (
	"errors"

	"Eventplanning.go/Api/db"
	"Eventplanning.go/Api/utils"
)


type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	
	hashPassword, err := utils.HasPassword(user.Password)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Email, hashPassword)

	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()	

	user.ID = userID
	return err

}

func (user *User) ValidateUser() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retrivedPassword string
	err := row.Scan(&user.ID, &retrivedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passordIsValid := utils.CheckPasswordHash(user.Password, retrivedPassword)

	if !passordIsValid {
		return errors.New("invalid credentials")
	}

	
	return nil
}