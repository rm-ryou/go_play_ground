package main

func fn(i, w int, arry []int) bool {
	if i == 0 {
		if w == 0 {
			return true
		} else {
			return false
		}
	}

	if fn(i-1, w, arry) {
		return true
	}

	if fn(i-1, w-arry[i-1], arry) {
		return true
	}

	return false
}
