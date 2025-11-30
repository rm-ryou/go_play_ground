package problem

func CountValue(a []int, key int) int {
	count := 0

	for _, v := range a {
		if v == key {
			count++
		}
	}

	return count
}
