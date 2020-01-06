package app

import (
	"net/http"

	"github.com/VIBHOR94/golang-microservices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
}
