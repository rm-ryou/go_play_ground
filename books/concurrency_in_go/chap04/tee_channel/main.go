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
				if ok == false {
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

func tee(done <-chan any, in <-chan any) (_, _ <-chan any) {
	out1, out2 := make(chan any), make(chan any)

	go func() {
		defer close(out1)
		defer close(out2)
		for val := range orDone(done, in) {
			var out1, out2 = out1, out2
			for range 2 {
				select {
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
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

func main() {
	done := make(chan any)
	defer close(done)

	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}
