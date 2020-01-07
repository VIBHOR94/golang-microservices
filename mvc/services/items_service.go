package services

import (
	"net/http"

	"github.com/VIBHOR94/golang-microservices/mvc/domain"
	"github.com/VIBHOR94/golang-microservices/mvc/utils"
)

type itemsService struct {
}

// ItemsService - Export variable ItemsService
var (
	ItemsService itemsService
)

// GetUser - Service to fetch the item
func (s *itemsService) GetItem(itemID int64) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
