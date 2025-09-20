package main

import "fmt"

func orDone(done, c <-chan any) <-chan any {
	valCh := make(chan any)
	go func() {
		defer close(valCh)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valCh <- v:
				case <-done:
				}
			}
		}
	}()
	return valCh
}

func bridge(done <-chan any, chanCh <-chan <-chan any) <-chan any {
	valCh := make(chan any)
	go func() {
		defer close(valCh)
		for {
			var ch <-chan any
			select {
			case maybeCh, ok := <-chanCh:
				if ok == false {
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
