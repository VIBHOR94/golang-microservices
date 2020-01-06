package controllers

import "net/http"

func GetUser(res http.ResponseWriter, req *http.Request) {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
