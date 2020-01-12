package services

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/VIBHOR94/golang-microservices/src/api/config"
	"github.com/VIBHOR94/golang-microservices/src/api/domain/github"
	"github.com/VIBHOR94/golang-microservices/src/api/domain/repositories"
	"github.com/VIBHOR94/golang-microservices/src/api/log"
	githubprovider "github.com/VIBHOR94/golang-microservices/src/api/providers/github_provider"
	"github.com/VIBHOR94/golang-microservices/src/api/utils/errors"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(clientID string, repositories repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
	CreateRepos(clientID string, repositories []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(clientID string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	log.Info("About to send request to extrnal API", fmt.Sprintf("client_id:%s", clientID), "status: pending")
	response, err := githubprovider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		log.Info("Recieved response from extrnal API", fmt.Sprintf("client_id:%s", clientID), "status: Error")
		return nil, errors.NewAPIError(err.StatusCode, err.Message)
	}
	log.Info("Recieved response from extrnal API", fmt.Sprintf("client_id:%s", clientID), "status: Success")

	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (s *reposService) CreateRepos(clientID string, requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)

	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	for _, current := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(clientID, current, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}

	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}
	return result, nil
}

func (s *reposService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse
	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (s *reposService) createRepoConcurrent(clientID string, input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	result, err := s.CreateRepo(clientID, input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	output <- repositories.CreateRepositoriesResult{Response: result}

}
