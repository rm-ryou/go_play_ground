package main

import "fmt"

func toString(done <-chan any, valueCh <-chan any) <-chan string {
	stringCh := make(chan string)
	go func() {
		defer close(stringCh)
		for v := range valueCh {
			select {
			case <-done:
				return
			case stringCh <- v.(string):
			}
		}
	}()
	return stringCh
}

func repeat(done <-chan any, values ...any) <-chan any {
	valueCh := make(chan any)
	go func() {
		defer close(valueCh)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueCh <- v:
				}
			}
		}
	}()
	return valueCh
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
