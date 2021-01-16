package main

import "github.com/gofiber/fiber/v2"

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Title     string `json:"title"`
}

type Users []User

func hello(c *fiber.Ctx) error {
	return c.JSON("Hello world")
}

func getUsers(c *fiber.Ctx) error {
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
	return c.JSON(u)
}

func main() {
	app := fiber.New()
	app.Get("/", hello)
	app.Get("/users", getUsers)
	app.Listen(":8080")
}
