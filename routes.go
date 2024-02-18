package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	//dbqueries
	insertUser  = `INSERT INTO users (username,password,email) VALUES (:UserName,:Password,:Email)`
	getUserId   = "Select userid from users where email = $1"
	getuserPass = "Select password from users where UserId = $1"
)

type User struct {
	UserId   int    `json:"userid"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type Response struct {
	Msg    string `json:"message"`
	UserId int    `json:"userid"`
}

func userLogin(c echo.Context) error {
	u := User{}
	var (
		hashPass string
		password string
	)

	c.Bind(&u)

	db.Get(&password, getuserPass, u.UserId)
	if err != nil {
		fmt.Println("user password err", err)
		return err
	}

	if ok := comparePasswords(hashPass, []byte(hashPass)); !ok {
		fmt.Println("password is incorrect")
		return errors.New("password is incorrect")
	}
	rsp := Response{
		Msg: " User Logged in",
	}
	return c.JSON(http.StatusOK, rsp)
}

func userSignUp(c echo.Context) error {
	u := User{}
	c.Bind(&u)
	p := hashAndSalt([]byte(u.Password))
	fmt.Println("Password	:	", p)
	fmt.Println("u.Username", u.UserName, u.Email)
	_, err = db.NamedExec(insertUser,
		map[string]interface{}{
			"UserName": u.UserName,
			"Password": p,
			"Email":    u.Email,
		})
	if err != nil {
		fmt.Println("Insert err :	", err)
	}
	user := User{}
	err = db.Get(&user, getUserId, u.Email)
	if err != nil {
		fmt.Println("getUserId err", err)
		return err
	}

	rsp := Response{
		Msg:    " User got created",
		UserId: user.UserId,
	}

	return c.JSON(http.StatusOK, rsp)
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
