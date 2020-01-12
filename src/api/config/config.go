package config

import "os"

const (
	secretGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	LogLevel                = "info"
)

var (
	githubAccessToken = os.Getenv(secretGithubAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}

func IsProduction() bool {
	return os.Getenv("GO_ENVIRONMENT") == "PRODUCTION"
}
