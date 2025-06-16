package dtos

type Track struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Link   string `json:"link"`
}

type Playlist struct {
	Name   string  `json:"name"`
	Tracks []Track `json:"tracks"`
}

type RecommendationResponse struct {
	BeerStyle string   `json:"beerStyle"`
	Playlist  Playlist `json:"playlist"`
}

type RecommendationRequest struct {
	Temperature float64 `json:"temperature"`
}
