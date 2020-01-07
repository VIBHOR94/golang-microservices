package domain

import (
	"fmt"
	"net/http"

	"github.com/VIBHOR94/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Fede", LastName: "Leon", Email: "myemail@gmail.com"},
	}

	// UserDao - Global level variabe of type userDao
	UserDao userDao
)

type userDao struct{}

// GetUser - Function that returns user by searching via userID
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exist", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
