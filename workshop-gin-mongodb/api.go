package demo

import (
	"demo/db"
	"demo/user"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// StartServer : Create new Router with Gin
func StartServer() {
	// gin.DisableConsoleColor()
	router := gin.Default()

	// Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.NoRoute(middlewares.NoRouteHandler())
	// router.NoMethod(middlewares.NoMethodHandler())

	// Prefix of all routes
	publicRoute := router.Group("/api/v1")

	// Initial resource from MongoDB
	resource, err := db.CreateResource()
	if err != nil {
		logrus.Error(err)
	}
	defer resource.Close()

	// Route of documents
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Add routes of users
	user.NewUserAPI(publicRoute, resource)

	// Start server
	router.Run() // listen and serve on 0.0.0.0:8080
}
