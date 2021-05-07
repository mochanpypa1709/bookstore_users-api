package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mochanpypa1709/bookstore_users-api/domain/users"
	"github.com/mochanpypa1709/bookstore_users-api/services"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		return
	}
	c.JSON(http.StatusCreated, result)
}
