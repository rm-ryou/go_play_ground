package main

import "fmt"

func swapPointer(i, j *int) (int, int) {
	tmp := i
	i = j
	j = tmp
	return *i, *j
}

func main() {
	x, y := 42, 2701
	p := &x
	fmt.Println(*p)

	*p = 21
	fmt.Println(x)

	p = &y
	*p = *p / 37
	fmt.Println(y)

	var pointer *int
	fmt.Printf("zero value of pointer: %v, Type: %T\n", pointer, pointer)

	i, j := 42, 24
	fmt.Println(swapPointer(&i, &j))
}
