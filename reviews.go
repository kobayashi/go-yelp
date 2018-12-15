package yelp

// Review is a review information for a business
type Review struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Text        string `json:"text"`
	Rating      int    `json:"rating"`
	TimeCreated string `json:"time_created"`
	User        struct {
		ProfilURL string `json:"profile_url"`
		ImageURL  string `json:"image_url"`
		Name      string `json:"name"`
	} `json:"user"`
}

// Reviews is a list of Review
type Reviews struct {
	Reviews           []Review `json:"reviews"`
	Total             int      `json:"total"`
	PossibleLanguages []string `json:"possible_languages"`
}
