package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	timeStart := time.Now()

	apis := []string{
		"https://pkg.go.dev",
		"https://go.dev",
		"https://api.github.com",
		"https://coursera.org",
		"https://en.wikipedia.org",
        "https://notworking.url",
	}

	channel := make(chan string)
	for _, api := range apis {
		go checkAPI(api, channel)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-channel)
	}

	elapsed := time.Since(timeStart)
	fmt.Printf("Execution Time: %s", elapsed)
}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("Error API %s is down!", api)
		return
	}

	ch <- fmt.Sprintf("API %s is Working!", api)
}
