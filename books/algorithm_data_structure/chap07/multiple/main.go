package main

func multiple(a, b []int, n int) int {
	sum := 0
	for i := n - 1; i >= 0; i-- {
		a[i] += sum
		remainder := a[i] % b[i]
		d := 0
		if remainder != 0 {
			d = b[i] - remainder
			sum += d
		}
	}
	return sum
}

func main() {}
