package main

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
				case ch <- v:
				case <-done:
				}
			}
		}
	}()

	return ch
}
