package yelp

import (
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

// func (e APIError) Error() error {
// 	if strings.Count(e.Code, "") > 0 || strings.Count(e.Description, "") > 0 {
// 		return fmt.Errorf("yelp: %s %s", e.Code, e.Description)
// 	}
// 	return nil
// }

func (e APIError) Error() string {
	return fmt.Sprintf("yelp: %s %s", e.Code, e.Description)
}
