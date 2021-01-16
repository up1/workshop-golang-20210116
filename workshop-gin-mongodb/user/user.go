package user

import (
	"demo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewUserAPI to create the router of user
func NewUserAPI(app *gin.RouterGroup, resource *db.Resource) {
	// Create repository
	repository := NewUserRepository(resource)

	app.GET("/users", handleGetUsers(repository))
	app.GET("/users/:id", handleGetUserByID(repository))
	app.POST("/users", handleCreateNewTask(repository))
}

type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Handlers

// GetUsers godoc
// @Summary Retrieves users based on query
// @Description Get Users
// @Produce json
// @Param name query string false "Name"
// @Param age query int false "Age"
// @Success 200 {array} Users
// @Router /api/v1/users [get]
func handleGetUsers(repository Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		code := http.StatusOK
		users, err := repository.GetAll()
		if err != nil {
			code = http.StatusInternalServerError
		}
		if len(users) == 0 {
			code = http.StatusNotFound
		}
		c.JSON(code, users)
	}
}

func handleGetUserByID(repository Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		code := http.StatusOK
		id := c.Param("id")
		user, err := repository.GetByID(id)
		response := map[string]interface{}{
			"user": user,
			"err":  getErrorMessage(err),
		}
		c.JSON(code, response)
	}
}

func handleCreateNewTask(repository Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		code := http.StatusOK
		newUser := UserRequest{}
		if err := c.Bind(&newUser); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Validate input !!!

		// Create data in database
		user, err := repository.CreateOne(newUser)
		response := map[string]interface{}{
			"user": user,
			"err":  getErrorMessage(err),
		}
		c.JSON(code, response)
	}
}

func getErrorMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
