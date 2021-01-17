package demo

import (
	"demo/db"
	"demo/user"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// StartServer : Create new Router with Gin
func StartServer() {
	router := echo.New()

	// ===== Middlewares

	// ===== Prefix of all routes
	publicRoute := router.Group("/api/v1")

	// ===== Initial resource from MongoDB
	resource, err := db.CreateResource()
	if err != nil {
		logrus.Error(err)
	}
	defer resource.Close()

	// ===== Add routes of users
	user.NewUserAPI(publicRoute, resource)

	// ===== Start server
	router.Start(":8080") // listen and serve on 0.0.0.0:8080
}
