package main

import "fmt"

func orDone(done, c <-chan any) <-chan any {
	ch := make(chan any)

	go func() {
		defer close(ch)

		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}

				select {
				case <-done:
				case ch <- v:
				}
			}
		}
	}()

	return ch
}

func bridge(done <-chan any, chanCh <-chan <-chan any) <-chan any {
	valCh := make(chan any)
	go func() {
		defer close(valCh)

		for {
			var ch <-chan any
			select {
			case <-done:
				return
			case maybeCh, ok := <-chanCh:
				if !ok {
					return
				}
				ch = maybeCh
			}

			for val := range orDone(done, ch) {
				select {
				case <-done:
				case valCh <- val:
				}
			}
		}
	}()

	return valCh
}

func genVals() <-chan <-chan any {
	chanCh := make(chan (<-chan any))
	go func() {
		defer close(chanCh)
		for i := range 10 {
			ch := make(chan any, 1)
			ch <- i
			close(ch)
			chanCh <- ch
		}
	}()
	return chanCh
}

func main() {
	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}
