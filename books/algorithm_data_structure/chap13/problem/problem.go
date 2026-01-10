package problem

type Point struct {
	X, Y int
}

type Queue struct {
	Data []Point
}

func (q *Queue) Enqueue(x, y int) {
	p := Point{X: x, Y: y}

	q.Data = append(q.Data, p)
}

func (q *Queue) Dequeue() Point {
	res := q.Data[0]
	q.Data = q.Data[1:]
	return res
}

func solveMaze(H, W int, maze []string) int {
	var sx, sy, gx, gy int
	for i := range maze {
		for j, v := range maze[i] {
			if v == 'S' {
				sx = j
				sy = i
			}

			if v == 'G' {
				gx = j
				gy = i
			}
		}
	}

	dist := make([][]int, H)
	for i := range dist {
		dist[i] = make([]int, W)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[sx][sy] = 0

	q := Queue{}
	q.Enqueue(sx, sy)

	// 		0
	// 	1		2
	// 		3
	dx := []int{0, -1, 1, 0}
	dy := []int{1, 0, 0, -1}
	for len(q.Data) != 0 {
		p := q.Dequeue()

		for dir := range 4 {
			nx := p.X + dx[dir]
			ny := p.Y + dy[dir]

			if nx < 0 || nx >= W || ny < 0 || ny >= H {
				continue
			}

			if maze[nx][ny] == '#' {
				continue
			}

			if dist[nx][ny] == -1 {
				dist[nx][ny] = dist[p.X][p.Y] + 1
				q.Enqueue(nx, ny)
			}
		}
	}

	return dist[gx][gy]
}
