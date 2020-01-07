package services

import (
	"github.com/VIBHOR94/golang-microservices/mvc/domain"
	"github.com/VIBHOR94/golang-microservices/mvc/utils"
)

type usersService struct {
}

// UsersService - Export variable UsersService
var (
	UsersService usersService
)

// GetUser - Service to fetch the user
func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userID)
}
