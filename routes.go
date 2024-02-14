package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	//dbqueries
	insertUser  = "INSERT INTO Users (Username, Password, Email) VALUES (?, ?, ?)"
	getUserId   = "Select UserID from Users where Email = ?"
	getuserPass = "Select Password from Users where UserId = ?"
)

type User struct {
	UserId   int    `json:"UserId"`
	UserName string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}

func userLogin(c echo.Context) error {
	u := User{}
    var hashPass string
	c.Bind(&u)
	row := db.QueryRow(getuserPass, u.UserId)
	err = row.Scan(&hashPass)
	if err != nil {
		fmt.Println("user password err", err)
		return err
	}

	if ok := comparePasswords(hashPass,[]byte(hashPass));!ok {
		fmt.Println("password is incorrect")
		return errors.New("password is incorrect")
	}
   fmt.Println("user loged in")
	
	return nil
}

func userSignUp(c echo.Context) error {
	u := User{}
	c.Bind(&u)
	p := hashAndSalt([]byte(u.Password))
	_, err := db.Exec(insertUser, u.UserName, p, u.Email)
	if err != nil {
		log.Fatal("err	:	", err)
		return err
	}

	row := db.QueryRow(getUserId, u.Email)
	err = row.Scan(&u.UserId)
	if err != nil {
		fmt.Println("getUserId err", err)
		return err
	}

	return nil
}

func userLogout(c echo.Context) error {

	return nil
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
