package main

import "fmt"

func toString(done <-chan any, valueCh <-chan any) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)

		for v := range valueCh {
			select {
			case <-done:
				return
			case ch <- v.(string):
			}
		}
	}()

	return ch
}

func repeat(done <-chan any, values ...any) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)

		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case ch <- v:
				}
			}
		}
	}()

	return ch
}

func take(done <-chan any, valueCh <-chan any, num int) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)

		for range num {
			select {
			case <-done:
				return
			case ch <- <-valueCh:
			}
		}
	}()

	return ch
}

func main() {
	done := make(chan any)
	defer close(done)

	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		message += token
	}

	fmt.Printf("message: %s...", message)
}
