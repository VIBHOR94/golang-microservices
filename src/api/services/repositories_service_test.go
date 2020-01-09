package services

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/VIBHOR94/golang-microservices/src/api/domain/repositories"

	"github.com/VIBHOR94/golang-microservices/src/api/clients/restclient"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputNmae(t *testing.T) {
	request := repositories.CreateRepoRequest{}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repository name", err.Message())
}
