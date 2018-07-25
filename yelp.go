package yelp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

const (
	baseURL = "https://api.yelp.com/v3/"
)

// Service provides a method for accessing Yelp endpoint
type Service struct {
	sling *sling.Sling
}

// Response is a typical response
type Response struct {
	Error    APIError        `json:"error"`
	Response json.RawMessage `json:"response"`
}

type APIError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client, apikey string) *Service {
	return &Service{
		sling: sling.New().Client(httpClient).Base(baseURL).Set("Authorization", "Bearer "+apikey),
	}
}

func (e APIError) Error() string {
	return fmt.Sprintf("yelp: %s %s", e.Code, e.Description)
}
