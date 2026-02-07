package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"tmdb-cli/models"
)

func FetchMovies(apiKey string) ([]models.Movie, error) {
	var data models.MovieResponse
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

func CheckApiKey(fileNameWApi string) string {
	var validApi models.ValidApi
	validApi.Success = true
	apiKey, err := os.ReadFile(fileNameWApi)
	if err != nil {
		fmt.Println("Error: TMDB API Key is missing!")
		fmt.Println("To fix this, run: go run main.go -key \"YOUR_API_KEY\"")

		return ""
	}
	url := "https://api.themoviedb.org/3/movie/popular?api_key=" + string(apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &validApi)
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}
	if !validApi.Success {
		fmt.Println(validApi.Message)
		return ""
	}

	return string(apiKey)

}

func SaveApiKey(fileNameWApi string, apiKey string) error {
	err := os.WriteFile(fileNameWApi, []byte(apiKey), 0644)
	if err != nil {
		panic(err)
	}
	return nil
}
