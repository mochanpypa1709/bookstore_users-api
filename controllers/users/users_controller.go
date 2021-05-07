package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mochanpypa1709/bookstore_users-api/domain/users"
	"github.com/mochanpypa1709/bookstore_users-api/services"
	"github.com/mochanpypa1709/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "Invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		return
	}
	c.JSON(http.StatusCreated, result)
}
