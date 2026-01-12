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

type Edge struct {
	to     int
	weight int
}

const INF int = 1e+010

func ScoreAttack(N, M int, x, y, w []int) int {
	graph := make([][]Edge, N)
	for i := range M {
		edge := Edge{
			to:     y[i] - 1,
			weight: -w[i],
		}
		graph[x[i]-1] = append(graph[x[i]-1], edge)
	}

	dist := make([]int, N)
	for i := range dist {
		dist[i] = INF
	}
	dist[0] = 0
	existNegativeCycle := make([]bool, N)

	for range N - 1 {
		for v := range N {
			for _, w := range graph[v] {
				val := dist[v] + w.weight
				if dist[w.to] > val {
					dist[w.to] = val
				}
			}
		}
	}

	for range N {
		for v := range N {
			for _, e := range graph[v] {
				val := dist[v] + e.weight
				if dist[e.to] > val {
					existNegativeCycle[e.to] = true
				}
				if existNegativeCycle[v] {
					existNegativeCycle[e.to] = true
				}
			}
		}
	}

	if !existNegativeCycle[N-1] {
		return -dist[N-1]
	}
	return -1
}
