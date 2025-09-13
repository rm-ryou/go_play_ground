package main

import "fmt"

type MyFloat float64

func do(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes  long\n", v, len(v))
	case MyFloat:
		fmt.Printf("%v is type of MyFloat\n", v)
		fmt.Printf("i: %v, v: %v\n", i, v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
		fmt.Printf("value: %v\n", i)
		if i == v {
			fmt.Println("true!!")
		}
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
	do(MyFloat(10))
}
