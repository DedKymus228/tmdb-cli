package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"tmdb-cli/api"
)

func main() {

	moviesType := flag.String("type", "popular", "Category of movies to fetch")
	apiKey := os.Getenv("API_KEY")
	endpoints := map[string]string{
		"popular":  "popular",
		"top":      "top_rated",
		"upcoming": "upcoming",
	}
	flag.Parse()
	category, ok := endpoints[*moviesType]
	if !ok {
		fmt.Println("no category found")
	}

	url := "https://api.themoviedb.org/3/movie/" + category + "?api_key=" + apiKey

	movieOUT, _ := api.FetchMovies(url)
	fmt.Println(url)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)
	for place, movie := range movieOUT {
		fmt.Fprintf(w, " #%d \t %s  \t Рейтинг: %.1f\n", place+1, movie.Title, movie.VoteAverage)
		//todo tabulation

	}
	w.Flush()

}
