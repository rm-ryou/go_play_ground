package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch false {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	case false:
		fmt.Println("false")
	default:
		fmt.Println("Good evening.")
	}
}
