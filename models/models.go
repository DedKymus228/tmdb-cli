package models

type Movie struct {
	Title       string  `json:"title"`
	VoteAverage float64 `json:"vote_average"`
	Overview    string  `json:"overview"`
	VoteCount   int     `json:"vote_count"`
	ReleaseDate string  `json:"release_date"`
	Adult       bool    `json:"adult"`
}

type MovieResponse struct {
	Results []Movie `json:"results"`
}

type ValidApi struct {
	Message string `json:"status_message"`
	Success bool   `json:"success"`
}
