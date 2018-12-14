package yelp

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type Business struct {
	ID           string      `json:"id"`
	Alias        string      `json:"alias"`
	Name         string      `json:"name"`
	ImageURL     string      `json:"image_url"`
	IsClosed     bool        `json:"is_closed"`
	URL          string      `json:"url"`
	ReviewCount  int         `json:"review_count"`
	Categories   Categories  `json:"categories"`
	Rating       int         `json:"rating"`
	Coordinates  Coordinates `json:"coordinates"`
	Transactions []string    `json:"transactions"`
	Price        string      `json:"price"`
	Location     Location    `json:"location"`
	Phone        string      `json:"phone"`
	DisplayPhone string      `json:"display_phone"`
	Distance     float64     `json:"distance"`
}

type Businesses struct {
	Total       int        `json:"total"`
	Bussinesses []Business `json:"businesses"`
	Region      Region     `json:"region"`
}

// BusinessService provides a method for accessing Yelp event endpoint
type BusinessService struct {
	sling *sling.Sling
}

type BusinessSearchParams struct {
	Term       string     `url:"term,omitempty"`
	Location   string     `url:"location,omitempty"`
	Latitude   float64    `url:"latitude,omitempty"`
	Longitude  float64    `url:"longitude,omitempty"`
	Radius     int        `url:"radius,omitempty"`
	Categories Categories `url:"categories,omitempty"`
	Locale     string     `url:"locale,omitempty"`
	Limit      int        `url:"limit,omitempty"`
	Offset     int        `url:"offset,omitempty"`
	SortBy     string     `url:"sort_by,omitempty"`
	Price      string     `url:"price,omitempty"`
	OpenNow    bool       `url:"open_now,omitempty"`
	OpenAt     int        `url:"open_at,omitempty"`
	Attributes string     `url:"attributes,omitempty"`
}

type BusinessMatchParams struct {
	Name           string  `url:"name,omitempty"`
	Address1       string  `url:"address_1,omitempty"`
	Address2       string  `url:"address_2,omitempty"`
	Address3       string  `url:"address_3,omitempty"`
	City           string  `url:"city,omitempty"`
	State          string  `url:"state,omitempty"`
	Country        string  `url:"country,omitempty"`
	Latitude       float64 `url:"latitude,omitempty"`
	Longitude      float64 `url:"longitude,omitempty"`
	Phone          string  `url:"phone,omitempty"`
	ZipCode        string  `url:"zip_code,omitempty"`
	YelpBusinessID string  `url:"yelp_business_id,omitempty"`
	Limit          int     `url:"limit,omitempty"`
	MatchThreshold string  `url:"match_threshold,omitempty"`
}

type businessResp struct {
	Business Business `json:"business"`
}

type businessSearchResp struct {
	BusinessesResp Businesses `json:"businesses"`
}

type businessMatchResp struct {
	BusinessList []Business `json:"business_list"`
}

type businessReviewResp struct {
	Reviews Reviews `json:"reviews"`
}

// Search returns up to 1000 businesses based on the provided search criteria
func (s *Service) BusinessSearch(params *BusinessSearchParams) ([]Business, *http.Response, error) {
	bsr := new(Businesses)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("businesses/search").QueryStruct(params).Receive(bsr, apiError)
	if err == nil {
		err = apiError
	}
	fmt.Println(*bsr)

	return bsr.Bussinesses, resp, err
}

// BusinessPhoneSearch returns a list of businesses based on the provided phone number.
func (s *Service) BusinessPhoneSearch(phone string) (Businesses, *http.Response, error) {
	businesses := new(businessSearchResp)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("businesses/search/"+phone).Receive(businesses, apiError)
	if err == nil {
		err = apiError
	}

	return businesses.BusinessesResp, resp, err
}

// BusinessDetail returns detailed business content.
func (s *Service) BusinessDetail(id string) (*Business, *http.Response, error) {
	business := new(Business)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("businesses/"+id).Receive(business, apiError)
	if err == nil {
		err = apiError
	}

	return business, resp, err
}

func (s *Service) BusinessMatch(params *BusinessMatchParams) ([]Business, *http.Response, error) {
	business := new(businessMatchResp)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("businesses/matches").QueryStruct(params).Receive(business, apiError)
	if err == nil {
		err = apiError
	}

	return business.BusinessList, resp, err
}

func (s *Service) BusinessReview(id string) (Reviews, *http.Response, error) {
	reviews := new(businessReviewResp)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("businesses/"+id+"reviews").Receive(reviews, apiError)
	if err == nil {
		err = apiError
	}

	return reviews.Reviews, resp, err
}
