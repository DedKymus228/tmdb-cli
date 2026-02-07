package main

import (
	"flag"
	"fmt"
	"tmdb-cli/api"
)

func main() {
	FileNameWApi := "api_key.txt"

	keyFlag := flag.String("key", "", "TMDB API Key")
	moviesType := flag.String("type", "popular", "Category of movies to fetch")
	endpoints := map[string]string{
		"playing":  "now_playing",
		"popular":  "popular",
		"top":      "top_rated",
		"upcoming": "upcoming",
	}
	flag.Parse()
	var apiKey string

	if *keyFlag != "" {
		api.SaveApiKey(FileNameWApi, *keyFlag)
		apiKey = api.CheckApiKey(FileNameWApi)
		return
	} else {
		apiKey = api.CheckApiKey(FileNameWApi)
	}
	category, ok := endpoints[*moviesType]
	if !ok {
		fmt.Println("no category found")
	}
	url := "https://api.themoviedb.org/3/movie/" + category + "?api_key=" + apiKey
	movieOUT, _ := api.FetchMovies(url)

	//fmt.Println(url)

	api.DisplayMovies(movieOUT)
}
