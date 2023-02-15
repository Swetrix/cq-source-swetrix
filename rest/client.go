package rest

import (
	"net/http"
)

type SwetrixRestClient struct {
	ProjectId string
	ApiKey    string

	HttpClient *http.Client
}

func NewSetrixClient(apiKey string) SwetrixRestClient {
	return SwetrixRestClient{
		ApiKey:     apiKey,
		HttpClient: &http.Client{},
	}
}
