package github

// ErrorResponse - To store error respose recieved while communiating to Github via api
type ErrorResponse struct {
	StatusCode       int     `json:"status_code"`
	Message          string  `json:"message"`
	DocumentationURL string  `json:"documentation_url"`
	Errors           []Error `json:"errors"`
}

// Error - A struct mentioning the errors occured. Part of the GithubErrorResponse strunct.
type Error struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
