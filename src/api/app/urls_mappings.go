package app

import (
	"github.com/VIBHOR94/golang-microservices/src/api/controllers/polo"
	"github.com/VIBHOR94/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
