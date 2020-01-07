package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/VIBHOR94/golang-microservices/mvc/utils"

	"github.com/VIBHOR94/golang-microservices/mvc/services"
)

// GetUser - Function to Get detais of user
func GetUser(res http.ResponseWriter, req *http.Request) {

	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		// Just return the bad request to the client
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonValue))
		return
	}

	user, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		// Handle the error and return to the client
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(apiErr.Message))
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
