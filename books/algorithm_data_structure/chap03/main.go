package main

func BasicLinearSearch(ary []int, key int) bool {
	flg := false

	for i := range ary {
		if i == key {
			flg = true
		}
	}

	return flg
}
