package yelp

import (
	"net/http"

	"github.com/dghubble/sling"
)

type Business struct {
	Rating       int         `json:"rating"`
	Price        string      `json:"price"`
	Phone        string      `json:"phone"`
	ID           string      `json:"id"`
	Alias        string      `json:"alias"`
	IsClosed     bool        `json:"is_closed"`
	Categories   Categories  `json:"categories"`
	ReviewCount  int         `json:"review_count"`
	Name         string      `json:"name"`
	URL          string      `json:"url"`
	Coordinates  Coordinates `json:"coordinates"`
	ImageURL     string      `json:"image_url"`
	Location     Location    `json:"location"`
	Distance     float64     `json:"distance"`
	Transactions []string    `json:"transactions"`
}

type Businesses struct {
	Total       int        `json:"total"`
	Bussinesses []Business `json:"bussinesses"`
	Region      Region     `json:"region"`
}

// BusinessService provides a method for accessing Yelp event endpoint
type BusinessService struct {
	sling *sling.Sling
}

type BusinessSearchParams struct {
	Term       string     `json:"term,omitempty"`
	Location   string     `json:"location,omitempty"`
	Latitude   float64    `json:"latitude,omitempty"`
	Longitude  float64    `json:"longitude,omitempty"`
	Radius     int        `json:"radius,omitempty"`
	Categories Categories `json:"categories,omitempty"`
	Locale     string     `json:"locale,omitempty"`
	Limit      int        `json:"limit,omitempty"`
	Offset     int        `json:"offset,omitempty"`
	SortBy     string     `json:"sort_by,omitempty"`
	Price      string     `json:"price,omitempty"`
	OpenNow    bool       `json:"open_now,omitempty"`
	OpenAt     int        `json:"open_at,omitempty"`
	Attributes string     `json:"attributes,omitempty"`
}

type BusinessMatchParams struct {
	Name           string  `json:"name,omitempty"`
	Address1       string  `json:"address_1,omitempty"`
	Address2       string  `json:"address_2,omitempty"`
	Address3       string  `json:"address_3,omitempty"`
	City           string  `json:"city,omitempty"`
	State          string  `json:"state,omitempty"`
	Country        string  `json:"country,omitempty"`
	Latitude       float64 `json:"latitude,omitempty"`
	Longitude      float64 `json:"longitude,omitempty"`
	Phone          string  `json:"phone,omitempty"`
	ZipCode        string  `json:"zip_code,omitempty"`
	YelpBusinessID string  `json:"yelp_business_id,omitempty"`
	Limit          int     `json:"limit,omitempty"`
	MatchThreshold string  `json:"match_threshold,omitempty"`
}

type businessResp struct {
	Business Business `json:"business"`
}

type businessSearchResp struct {
	Businesses Businesses `json:"businesses"`
}

type businessMatchResp struct {
	BusinessList []Business `json:"business_list"`
}

type businessReviewResp struct {
	Reviews Reviews `json:"reviews"`
}

// Search returns up to 1000 businesses based on the provided search criteria
func (s *Service) BusinessSearch(params *BusinessSearchParams) (Businesses, *http.Response, error) {
	businesses := new(businessSearchResp)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("businesses/search").QueryStruct(params).Receive(businesses, apiError)
	if err == nil {
		err = apiError
	}

	return businesses.Businesses, resp, err
}

// BusinessPhoneSearch returns a list of businesses based on the provided phone number.
func (s *Service) BusinessPhoneSearch(phone string) (Businesses, *http.Response, error) {
	businesses := new(businessSearchResp)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("businesses/search/"+phone).Receive(businesses, apiError)
	if err == nil {
		err = apiError
	}

	return businesses.Businesses, resp, err
}

// BusinessDetail returns detailed business content.
func (s *Service) BusinessDetail(id string) (Business, *http.Response, error) {
	business := new(businessResp)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("businesses/"+id).Receive(business, apiError)
	if err == nil {
		err = apiError
	}

	return business.Business, resp, err
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
