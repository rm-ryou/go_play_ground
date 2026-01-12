package problem

func calcLongestPath(graph [][]int, dp []int, cur int) int {
	if dp[cur] != -1 {
		return dp[cur]
	}

	var res int
	for _, v := range graph[cur] {
		b := calcLongestPath(graph, dp, v) + 1
		if res < b {
			res = b
		}
	}

	dp[cur] = res
	return dp[cur]
}

func LongestPath(N, M int, x, y []int) int {
	graph := make([][]int, N)
	for i := range M {
		graph[x[i]-1] = append(graph[x[i]-1], y[i]-1)
	}

	dp := make([]int, N*N)
	for i := range dp {
		dp[i] = -1
	}

	res := 0
	for i := range N {
		val := calcLongestPath(graph, dp, i)
		if res < val {
			res = val
		}
	}

	return res
}
