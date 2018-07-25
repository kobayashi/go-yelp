package yelp

// Review is a review information for a business
type Review struct {
	ID     string `json:"id"`
	Rating int    `json:"rating"`
	User   struct {
		ImageURL string `json:"image_url"`
		Name     string `json:"name"`
	} `json:"user"`
	Text        string `json:"text"`
	TimeCreated string `json:"time_created"`
	URL         string `json:"url"`
}

// Reviews is a list of Review
type Reviews struct {
	Reviews           []Review
	Total             int
	PossibleLanguages []string
}
