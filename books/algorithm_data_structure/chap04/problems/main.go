package main

var memo []int

func tribonacciMemo(n int) int {
	if n == 0 || n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	return memo[n]
}

func tribonacci(n int) int {
	if n == 0 || n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
}

func fn755(n, curNum, use int, counter *int) {
	if curNum >= n {
		return
	}
	if use == 0x111 {
		*counter++
	}

	fn755(n, curNum*10+7, use|0x100, counter)
	fn755(n, curNum*10+5, use|0x010, counter)
	fn755(n, curNum*10+3, use|0x001, counter)

}

var memoForSumOfPart [][]int

func sumOfPart(n, w int, arry []int) int {
	if n == 0 {
		if w == 0 {
			return 1
		} else {
			return 0
		}
	}

	if memoForSumOfPart[n][w] != 0 {
		return memoForSumOfPart[n][w]
	}
	if sumOfPart(n-1, w, arry) == 1 {
		memoForSumOfPart[n][w] = 1
		return 1
	}

	if sumOfPart(n-1, w-arry[n-1], arry) == 1 {
		memoForSumOfPart[n][w] = 1
		return 1
	}

	return memoForSumOfPart[n][w]
}

func main() {
}
