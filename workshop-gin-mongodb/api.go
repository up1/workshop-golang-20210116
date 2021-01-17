package demo

import (
	"demo/db"
	"demo/middlewares"
	"demo/user"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/penglongli/gin-metrics/ginmetrics"
)

// StartServer : Create new Router with Gin
func StartServer() {
	// gin.DisableConsoleColor()
	// router := gin.Default()
	router := gin.New()

	// ===== Middlewares
	// router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(middlewares.AuthRequired())
	router.NoRoute(middlewares.NoRouteHandler())
	// router.NoMethod(middlewares.NoMethodHandler())

	// ===== Prometheus
	// get global Monitor object
	m := ginmetrics.GetMonitor()
	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	// set middleware for gin
	m.Use(router)

	// ===== Prefix of all routes
	publicRoute := router.Group("/api/v1")

	// ===== Initial resource from MongoDB
	resource, err := db.CreateResource()
	if err != nil {
		logrus.Error(err)
	}
	defer resource.Close()

	// ===== Route of documents
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ===== Add routes of users
	user.NewUserAPI(publicRoute, resource)

	// ===== Start server
	router.Run() // listen and serve on 0.0.0.0:8080
}
