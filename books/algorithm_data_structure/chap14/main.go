package main

const INF int = 1e+06

type Edge struct {
	to int
	w  int
}

func BellmanFord(N, M, s, v int, a, b, w []int) int {
	graph := make([][]Edge, M)
	for i := range graph {
		graph[a[i]] = append(graph[a[i]], Edge{to: b[i], w: w[i]})
	}

	existNegativeCycle := false
	dist := make([]int, N)
	for i := range dist {
		dist[i] = INF
	}
	dist[s] = 0

	for i := range N {
		update := false
		for j := range N {
			if dist[j] == INF {
				continue
			}

			for _, e := range graph[j] {
				a := dist[e.to]
				b := dist[j] + e.w
				if a > b {
					dist[e.to] = b
					update = true
				}

			}
		}

		if !update {
			break
		}

		if i == N-1 && update {
			existNegativeCycle = true
		}
	}

	if existNegativeCycle {
		return -1
	}

	return dist[v]
}

func Dijkstr(N, M, s, t int, a, b, w []int) int {
	graph := make([][]Edge, M)
	for i := range graph {
		graph[a[i]] = append(graph[a[i]], Edge{to: b[i], w: w[i]})
	}

	used := make([]bool, N)
	dist := make([]int, N)
	for i := range dist {
		dist[i] = INF
	}
	dist[s] = 0

	for range N {
		minDist := INF
		minVal := -1

		for v := range N {
			if !used[v] && dist[v] < minDist {
				minDist = dist[v]
				minVal = v
			}
		}
		if minVal == -1 {
			break
		}

		for _, e := range graph[minVal] {
			a := dist[e.to]
			b := dist[minVal] + e.w

			if a > b {
				dist[e.to] = b
			}
		}
		used[minVal] = true
	}

	return dist[t]
}
