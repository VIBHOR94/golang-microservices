package services

import (
	"github.com/VIBHOR94/golang-microservices/mvc/domain"
)

// GetUser - Service to fetch the user
func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}
