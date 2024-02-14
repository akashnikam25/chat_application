package main

import (
	"github.com/labstack/echo/v4"
)

const (
	//dbqueries 
	insertUser = "INSERT INTO users (Username, Password, Email) VALUES (?, ?, ?)"
	getUserId  = "Select * from user"
)
type User struct {
	UserId   int    `json:"UserId"`
	UserName string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}

func userLogin(c echo.Context) error {

	return nil
}

func userSignUp(c echo.Context) error {
	u := User{}
	c.Bind(&u)
	
	db.Exec(insertUser,u.UserName,u.Password,u.Email)
	

	return nil
}

func userLogout(c echo.Context) error {

	return nil
}
