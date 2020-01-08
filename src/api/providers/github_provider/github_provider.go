package githubprovider

import (
	"fmt"

	"github.com/VIBHOR94/golang-microservices/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

// CreateRepo - Handles the Creation of gihub repository vai apis
func CreateRepo(accessToken string, request github.CreateRepoRequest) (github.CreateRepoResponse, github.ErrorResponse) {
	// headers := http.Header{}

	return github.CreateRepoResponse{}, github.ErrorResponse{}
}
