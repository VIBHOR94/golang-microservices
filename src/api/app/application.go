package app

import (
	"github.com/VIBHOR94/golang-microservices/src/api/log"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	log.Info("about to map the urls")
	mapUrls()
	log.Info("urls successfully mapped")
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
