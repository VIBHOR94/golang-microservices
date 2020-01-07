package utils

import "github.com/gin-gonic/gin"

// Respond - Function to convert returned type based on header
func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}

// RespondError - Function to take care of error handling
func RespondError(c *gin.Context, err *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, err)
		return
	}
	c.JSON(err.StatusCode, err)
}
