package yelp

import (
	"net/http"
	"time"
)

// Event is an event information.
type Event struct {
	AttendingCount  int      `json:"attending_count"`
	Category        string   `json:"category"`
	Cost            float64  `json:"cost"`
	CostMax         float64  `json:"cost_max"`
	Description     string   `json:"description"`
	EventSiteURL    string   `json:"event_site_url"`
	ID              string   `json:"id"`
	ImageURL        string   `json:"image_url"`
	InterestedCount int      `json:"interested_count"`
	IsCanceled      bool     `json:"is_canceled"`
	IsFree          bool     `json:"is_free"`
	IsOfficial      bool     `json:"is_official"`
	Latitude        float64  `json:"latitude"`
	Longitude       float64  `json:"longitude"`
	Name            string   `json:"name"`
	TicketsURL      string   `json:"tickets_url"`
	TimeEnd         time.TIME   `json:"time_end"`
	TimeStart       time.TIME   `json:"time_start"`
	Location        Location `json:"location"`
	BusinessID      string   `json:"business_id"`
}

type EventSearchParams struct {
	Locale         string   `url:"locale,omitempty"`
	Offset         int      `url:"offset,omitempty"`
	Limit          int      `url:"limit,omitempty"`
	SortBy         string   `url:"sort_by,omitempty"`
	SortOn         string   `url:"sort_on,omitempty"`
	StartDate      int      `url:"start_date,omitempty"`
	EndDate        int      `url:"end_date,omitempty"`
	Categories     string   `url:"categories,omitempty"`
	IsFree         bool     `url:"is_free,omitempty"`
	Location       string   `url:"location,omitempty"`
	Latitude       float64  `url:"latitude,omitempty"`
	Longitude      float64  `url:"longitude,omitempty"`
	Radius         int      `url:"radius,omitempty"`
	ExcludedEvents []string `url:"excluded_events,omitempty"`
}

type EventFeaturedParams struct {
	Locale    string  `url:"locale,omitempty"`
	Location  string  `url:"location,omitempty"`
	Latitude  float64 `url:"latitude,omitempty"`
	Longitude float64 `url:"longitude,omitempty"`
}

type eventResp struct {
	Event Event `json:"event"`
}

type eventsResp struct {
	Events []Event `json:"events"`
	Total  int     `json:"total"`
}

// EventSearch returns events based on the provided search criteria.
func (s *Service) EventSearch(params *EventSearchParams) ([]Event, *http.Response, error) {
	events := new(eventsResp)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("events").QueryStruct(params).Receive(events, apiError)
	if err == nil {
		err = apiError
	}

	return events.Events, resp, err
}

// EventLookUp returns the detailed information of a Yelp event
func (s *Service) EventLookUp(id string) (*Event, *http.Response, error) {
	event := new(Event)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("events/"+id).Receive(event, apiError)
	if err == nil {
		err = apiError
	}

	return event, resp, err
}

// EventFeatured returns the featured event for a given location
func (s *Service) EventFeatured(params *EventFeaturedParams) (*Event, *http.Response, error) {
	event := new(Event)
	apiError := new(APIError)

	resp, err := s.sling.New().Get("events/featured").QueryStruct(params).Receive(event, apiError)
	if err == nil {
		err = apiError
	}

	return event, resp, err
}
