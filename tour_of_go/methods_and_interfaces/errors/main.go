package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (me MyError) String() string {
	return "implement stringer"
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		When: time.Now(),
		What: "it din't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
