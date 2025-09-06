package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	customRand := rand.New(rand.NewSource(10))
	fmt.Println("My hate number is ", customRand.Intn(10))
}
