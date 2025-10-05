package main

import "fmt"

func orDone(done, valCh <-chan any) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case v, ok := <-valCh:
				if !ok {
					return
				}
				select {
				case ch <- v:
				case <-done:
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
			case maybeCh, ok := <-chanCh:
				if !ok {
					return
				}
				ch = maybeCh
			case <-done:
				return
			}

			for val := range orDone(done, ch) {
				select {
				case valCh <- val:
				case <-done:
				}
			}
		}
	}()
	return valCh
}

func main() {
	genVals := func() <-chan <-chan any {
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

	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}
