package main

func BasicLinearSearch(ary []int, key int) bool {
	flg := false

	for _, v := range ary {
		if v == key {
			flg = true
		}
	}

	return flg
}

func FindIndex(ary []int, key int) int {
	idx := -1

	for i, v := range ary {
		if v == key {
			idx = i
		}
	}

	return idx
}
