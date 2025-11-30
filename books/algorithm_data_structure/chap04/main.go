package main

func SumRecv(n int) int {
	if n == 0 {
		return 0
	}

	return n + SumRecv(n-1)
}

func GCD(m, n int) int {
	if n == 0 {
		return m
	}
	return GCD(n, m%n)
}

func FiboRecv(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}

	return FiboRecv(n-1) + FiboRecv(n-2)
}

func FiboMemo(memo []int, n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}

	if memo[n] != -1 {
		return memo[n]
	}

	memo[n] = FiboMemo(memo, n-1) + FiboMemo(memo, n-2)
	return memo[n]
}

func IsCreateValueFromAry(n, w int, ary []int) bool {
	if n == 0 {
		if w == 0 {
			return true
		} else {
			return false
		}
	}

	if IsCreateValueFromAry(n-1, w-ary[n-1], ary) {
		return true
	}

	if IsCreateValueFromAry(n-1, w, ary) {
		return true
	}

	return false
}
