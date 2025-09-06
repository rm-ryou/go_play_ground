package main

import "fmt"

func main() {
	sum := 1
	for {
		if sum > 1000 {
			break
		}
		sum += sum
	}
	fmt.Println(sum)
}
