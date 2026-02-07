package api

import (
	"fmt"
	"io"
	"net/http"
)

func PrintRawAPI(apiKey string) ([]byte, error) {

	resp, err := http.Get(apiKey)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return body, nil
}
