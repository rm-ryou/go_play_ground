package problem

func Tribo(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 0
	case 2:
		return 1
	}

	return Tribo(n-1) + Tribo(n-2) + Tribo(n-3)
}

func TriboMemo(memo []int, n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 0
	case 2:
		return 1
	}

	if memo[n] != -1 {
		return memo[n]
	}

	memo[n] = TriboMemo(memo, n-1) + TriboMemo(memo, n-2) + TriboMemo(memo, n-3)
	return memo[n]
}
