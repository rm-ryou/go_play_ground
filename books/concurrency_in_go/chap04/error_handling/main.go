package main

import (
	"fmt"
	"net/http"
)

type Response struct {
	Error    error
	response *http.Response
}

func checkStatus(done <-chan any, urls ...string) <-chan Response {
	responses := make(chan Response)
	go func() {
		defer close(responses)
		for _, url := range urls {
			res, err := http.Get(url)
			r := Response{Error: err, response: res}

			select {
			case <-done:
				return
			case responses <- r:
			}
		}
	}()
	return responses
}

func main() {
	done := make(chan any)
	defer close(done)

	for res := range checkStatus(done, []string{"https://www.google.com", "https://badhost"}...) {
		if res.Error != nil {
			fmt.Printf("error: %v\n", res.Error)
			continue
		}
		fmt.Printf("Response: %v\n", res.response.Status)
	}
}
