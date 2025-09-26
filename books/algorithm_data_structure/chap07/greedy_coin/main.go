package main

var values = []int{500, 100, 50, 10, 5, 1}

func greedyCoin(arry []int, x int) int {
	res := 0

	for i := range x {
		add := x / values[i]
		if add > arry[i] {
			add = arry[i]
		}

		x -= arry[i] * add
		res += add
	}
	return res
}

func main() {}
