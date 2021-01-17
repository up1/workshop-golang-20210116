package user

import (
	"demo/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

// NewUserAPI to create the router of user
func NewUserAPI(app *echo.Group, resource *db.Resource) {
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
func handleGetUsers(repository Repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		code := http.StatusOK
		users, err := repository.GetAll()
		if err != nil {
			code = http.StatusInternalServerError
		}
		if len(users) == 0 {
			code = http.StatusNotFound
		}
		return c.JSON(code, users)
	}
}

func handleGetUserByID(repository Repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		code := http.StatusOK
		id := c.Param("id")
		user, err := repository.GetByID(id)
		response := map[string]interface{}{
			"user": user,
			"err":  getErrorMessage(err),
		}
		return c.JSON(code, response)
	}
}

func handleCreateNewTask(repository Repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		code := http.StatusOK
		newUser := UserRequest{}
		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		// Validate input !!!

		// Create data in database
		user, err := repository.CreateOne(newUser)
		response := map[string]interface{}{
			"user": user,
			"err":  getErrorMessage(err),
		}
		return c.JSON(code, response)
	}
}

func getErrorMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
