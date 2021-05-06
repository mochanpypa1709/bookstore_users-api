package app

import "github.com/mochanpypa1709/bookstore_users-api/controllers"

func MapUrls() {
	router.GET("/ping", controllers.Ping)
}
