package main

func dfs(isReached []bool, graph [][]int, v int) {
	isReached[v] = true

	for _, v := range graph[v] {
		if isReached[v] {
			continue
		}

		dfs(isReached, graph, v)
	}
}

func IsSTPathExists(N, M, s, t int, a, b []int) bool {
	graph := make([][]int, N)
	for i := range M {
		graph[a[i]] = append(graph[a[i]], b[i])
	}

	isReached := make([]bool, N)
	dfs(isReached, graph, s)

	return isReached[t]
}
