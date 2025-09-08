package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
	Num   = 1<<64 - 1
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Println(Small)
	fmt.Println(Num >> 30)
	fmt.Printf("Value: %v, Type: %T\n", Num>>1, Num>>1)
	fmt.Printf("Value: %v, Type: %T\n", Num>>60, Num>>60)
}
