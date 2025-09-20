package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func checkStatus(done <-chan any, urls ...string) <-chan Result {
	results := make(chan Result)
	go func() {
		defer close(results)

		for _, url := range urls {
			var res Result
			resp, err := http.Get(url)
			res = Result{Error: err, Response: resp}
			select {
			case <-done:
				return
			case results <- res:
			}
		}
	}()
	return results
}

func main() {
	done := make(chan any)
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for res := range checkStatus(done, urls...) {
		if res.Error != nil {
			fmt.Printf("error: %v", res.Error)
			continue
		}
		fmt.Printf("Response: %v\n", res.Response.Status)
	}
}
