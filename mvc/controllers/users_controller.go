package controllers

import (
	"log"
	"net/http"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")
	log.Printf("about to process user with user id = %v", userID)
}
