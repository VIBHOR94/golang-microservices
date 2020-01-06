package services

import (
	"github.com/VIBHOR94/golang-microservices/mvc/domain"
	"github.com/VIBHOR94/golang-microservices/mvc/utils"
)

// GetUser - Service to fetch the user
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
