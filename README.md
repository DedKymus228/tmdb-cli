# TMDB-CLI-TOOL

A simple command-line interface (CLI) tool for fetching movie information from the TMDB API, written in Go. 

## Features

* Fetch and display Now Playing movies
* Fetch and display Popular movies
* Fetch and display Top Rated movies
* Fetch and display Upcoming movies

## Getting Started & Usage

```bash
# 1. Clone the repository
git clone [https://github.com/YOUR_USERNAME/tmdb-cli.git](https://github.com/YOUR_USERNAME/tmdb-cli.git)
cd tmdb-cli

# 2. Build the application
go build -o tmdb-cli (tmdb-cli.exe for Windows)

# 3. Save your API key (First run only)
./tmdb-cli -key "YOUR_TMDB_API_KEY"

# 4. Fetch different categories
./tmdb-cli -type popular        # Get popular movies (default)
./tmdb-cli -type top            # Get top rated movies
./tmdb-cli -type upcoming       # Get upcoming movies
./tmdb-cli -type playing        # Get movies currently in theaters

# 5. Using go run (Alternative)
go run main.go -type popular