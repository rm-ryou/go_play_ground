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

func Problem755(k, use, cur int, counter *int) {
	if cur > k {
		return
	}
	if use == 0b111 {
		*counter++
	}

	Problem755(k, use|0b001, cur*10+3, counter)
	Problem755(k, use|0b010, cur*10+5, counter)
	Problem755(k, use|0b100, cur*10+7, counter)
}
