package domain

import (
	"fmt"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Fede", LastName: "Leon", Email: "myemail@gmail.com"},
	}
)

// GetUser - Function that returns user by searching via userID
func GetUser(userID int64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, fmt.Errorf("user %d was not found", userID)
}
