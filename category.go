package yelp

import (
	"net/http"
)

type Category struct {
	Alias            string        `json:"alias,omitempty"`
	Title            string        `json:"title,omitempty"`
	ParentAliases    []string      `json:"parent_aliases,omitempty"`
	CountryWhitelist []interface{} `json:"country_whitelist,omitempty"`
	CountryBlacklist []interface{} `json:"country_blacklist,omitempty"`
}

type Categories []Category

type categoriesParams struct {
	Locale string `json:"locale,omitempty"`
}

type categoryDetailParams struct {
	Locale string `json:"locale,omitempty"`
	Alias  string `json:"alias"`
}

type categoriesResp struct {
	Categories Categories `json:"categories"`
}

type categoryResp struct {
	Category Category `json:"category"`
}

func (s *Service) Categories(params *categoriesParams) (Categories, *http.Response, error) {
	categories := new(categoriesResp)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("categories").QueryStruct(params).Receive(categories, apiError)
	if err == nil {
		err = apiError
	}

	return categories.Categories, resp, err
}

func (s *Service) CategoryDetail(alias string) (Category, *http.Response, error) {
	category := new(categoryResp)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("categories/"+alias).Receive(category, apiError)
	if err == nil {
		err = apiError
	}

	return category.Category, resp, err
}
