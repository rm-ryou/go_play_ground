package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intCh := make(chan int, 4)
	go func() {
		defer close(intCh)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := range 5 {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intCh <- i
		}
	}()

	for i := range intCh {
		fmt.Fprintf(&stdoutBuff, "Received %v\n", i)
	}
}
