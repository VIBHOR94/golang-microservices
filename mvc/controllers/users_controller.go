package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/VIBHOR94/golang-microservices/mvc/services"
)

// GetUser - Function to Get detais of user
func GetUser(res http.ResponseWriter, req *http.Request) {

	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		// Just return the bad request to the client
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("user_id must be a number"))
		return
	}

	user, err := services.GetUser(userID)
	if err != nil {
		// Handle the error and return to the client
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
