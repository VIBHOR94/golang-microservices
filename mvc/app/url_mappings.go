package app

import (
	"github.com/VIBHOR94/golang-microservices/mvc/controllers"
)

func mapURLs() {
	router.GET("/users/:user_id", controllers.GetUser)
}
