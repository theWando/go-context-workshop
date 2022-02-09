package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://xkcd.com/info.0.json", nil)
	if err != nil {
		fmt.Println("failed to create request", err)
		os.Exit(1)
	}

	client := &http.Client{
		Timeout: time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("failed to do request", err)
		os.Exit(1)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var c Comic
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		fmt.Println("failed to decode response", err)
		os.Exit(1)
	}
	fmt.Printf("latest comic is:\n%#v", c)
}
