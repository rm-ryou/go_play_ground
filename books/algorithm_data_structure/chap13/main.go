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

func calcIsBipartiteGraph(graph [][]int, color []int, v, cur int) bool {
	color[v] = cur

	for _, nextV := range graph[v] {
		if color[nextV] != -1 {
			if color[nextV] == cur {
				return false
			}

			continue
		}

		if !calcIsBipartiteGraph(graph, color, nextV, 1-cur) {
			return false
		}
	}

	return true
}

func IsBipartiteGraph(n, m int, a, b []int) bool {
	graph := make([][]int, n)
	for i := range m {
		graph[a[i]] = append(graph[a[i]], b[i])
		graph[b[i]] = append(graph[b[i]], a[i])
	}

	color := make([]int, n)
	for i := range color {
		color[i] = -1
	}

	res := true
	for i := range n {
		if color[i] != -1 {
			continue
		}
		if !calcIsBipartiteGraph(graph, color, i, 0) {
			res = false
		}
	}

	return res
}
