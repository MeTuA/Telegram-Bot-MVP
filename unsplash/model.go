package unsplash

type Credentials struct {
	ClientID string
}

// struct for response from Unsplash API
type RandomPhoto struct {
	Description string `json:"description"`
	URL         struct {
		Regular string `json:"regular"`
	} `json:"urls"`
}
