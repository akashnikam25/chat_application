package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

var (
	db  *sqlx.DB
	err error
)

func initDb() {
	db, err = sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=akash dbname=chatapp sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database:", err)
	}
	err = db.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL database!")
}
func main() {
	initDb()
	defer db.DB.Close()

	e := echo.New()
	// Use the CORS middleware to enable Cross-Origin Resource Sharing
	e.Use(middleware.CORS())
	e.POST("/user/signup", userSignUp)
	e.POST("/user/login", userLogin)
	e.POST("/user/logout", userLogout)
	e.POST("/user/coversation/", createConversation)

	// Start the Echo server
	fmt.Printf("Server is listening on port %s...\n", ":8000")
	e.Logger.Fatal(e.Start(":8000"))

}
