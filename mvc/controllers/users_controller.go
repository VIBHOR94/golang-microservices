package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/VIBHOR94/golang-microservices/mvc/utils"

	"github.com/VIBHOR94/golang-microservices/mvc/services"
)

// GetUser - Function to Get detais of user
func GetUser(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		// Just return the bad request to the client
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		// Handle the error and return to the client
		utils.RespondError(c, apiErr)
		return
	}

	// return user to client
	utils.Respond(c, http.StatusOK, user)
}
