package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	db  *sql.DB
	err error
)

func initDb() {
	conStr := "root:akash@tcp(127.0.0.1:3306)/chatapp1"
	db, err = sql.Open("mysql", conStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL database!")
}
func main() {
	defer db.Close()
	initDb()
	e := echo.New()
	// Use the CORS middleware to enable Cross-Origin Resource Sharing
	e.Use(middleware.CORS())
	e.POST("/user/signup", userSignUp)
	e.POST("/user/login", userLogin)
	e.POST("/user/logout", userLogout)

	// Start the Echo server
	e.Logger.Fatal(e.Start(":8000"))

}
