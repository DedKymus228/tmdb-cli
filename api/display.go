package api

import (
	"fmt"
	"os"
	"text/tabwriter"
	"tmdb-cli/models"
)

func DisplayMovies(movies []models.Movie) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "TITLE\tRELEASE\tRATING\tVOTES\tADULT")
	fmt.Fprintln(w, "-----\t-------\t------\t-----\t-----")

	for _, m := range movies {
		// Форматируем вывод
		fmt.Fprintf(w, "%s\t%s\t%.1f\t%d\t%v\n",
			m.Title,
			m.ReleaseDate,
			m.VoteAverage,
			m.VoteCount,
			m.Adult,
		)
	}

	w.Flush() // Выталкиваем данные из буфера в консоль
}
