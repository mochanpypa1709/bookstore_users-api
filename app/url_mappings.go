package app

import (
	"github.com/mochanpypa1709/bookstore_users-api/controllers/ping"
	"github.com/mochanpypa1709/bookstore_users-api/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
}
