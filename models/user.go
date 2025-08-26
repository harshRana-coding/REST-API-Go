package models

import (
	"rest-api/db"
)

type User struct {
	ID       int64
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users (email,password)
	VALUES (?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}
