package app

import "github.com/mochanpypa1709/bookstore_users-api/controllers"

func MapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/users", controllers.GetUser)
	router.GET("/users/:user_id", controllers.GetUser)
}
