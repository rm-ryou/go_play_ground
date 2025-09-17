package main

import (
	"fmt"
	"sync"
	"time"
)

type Button struct {
	Clicked *sync.Cond
}

var button = Button{
	Clicked: sync.NewCond(&sync.Mutex{}),
}

func subscribe(c *sync.Cond, fn func()) {
	var goroutineRunning sync.WaitGroup
	goroutineRunning.Add(1)
	go func() {
		goroutineRunning.Done()
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()
		fn()
	}()
	goroutineRunning.Wait()
}

func main() {
	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	time.Sleep(5 * time.Second)
	button.Clicked.Broadcast()
	clickRegistered.Wait()
}
