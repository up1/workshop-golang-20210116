package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Title     string `json:"title"`
}

type Users []User

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello world")
}

func getUser(c echo.Context) error {
	u := Users{
		User{
			Firstname: "f1",
			Lastname:  "l1",
			Title:     "Mr.",
		},
		User{
			Firstname: "f2",
			Lastname:  "l2",
			Title:     "Miss.",
		},
	}
	return c.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.GET("/users", getUser)
	e.Logger.Fatal(e.Start(":8080"))
}
