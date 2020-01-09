package services

import (
	"strings"

	githubprovider "github.com/VIBHOR94/golang-microservices/src/api/providers/github_provider"

	"github.com/VIBHOR94/golang-microservices/src/api/config"
	"github.com/VIBHOR94/golang-microservices/src/api/domain/github"

	"github.com/VIBHOR94/golang-microservices/src/api/domain/repositories"
	"github.com/VIBHOR94/golang-microservices/src/api/utils/errors"
)

type resposService struct{}

type reposServiceInterface interface {
	CreateRepo(repositories repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &resposService{}
}

func (s *resposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := githubprovider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewAPIError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
