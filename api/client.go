package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Movie struct {
	Title       string  `json:"title"`
	VoteAverage float64 `json:"vote_average"`
	Overview    string  `json:"overview"`
}

type MovieResponse struct {
	Results []Movie `json:"results"`
}

func FetchMovies(apiKey string) ([]Movie, error) {
	var data MovieResponse
	resp, err := http.Get(apiKey)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}
	return data.Results, nil
}
