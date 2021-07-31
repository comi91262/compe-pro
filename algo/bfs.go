package algo

// ただ BFSするだけ
func bfs(start int, g [][]int) {
	q := queue{}
	q.push(start)

	for !q.empty() {
		node := q.pop()
		for _, e := range g[node] {
			q.push(e)
		}
	}
}

// 単一コストの最短経路問題を BFS でとく
func bfsMinDist(start int, g [][]int) []int {
	const inf = -(1 << 60)
	q := queue{}
	q.push(start)

	dist := make([]int, len(g)+1)

	for i := 0; i <= len(g); i++ {
		dist[i] = inf
	}
	dist[start] = 0
	for !q.empty() {
		node := q.pop()
		for _, e := range g[node] {
			if dist[e] == inf {
				dist[e] = dist[node] + 1
				q.push(e)
			}
		}
	}

	return dist
}

// 単一コストの最短経路問題を BFS でとく (二次元版)
func bfs2MinDist(h, w int, sx, sy, ex, ey int) [][]int {
	const inf = 1 << 60
	var dx = [4]int{1, 0, -1, 0}
	var dy = [4]int{0, 1, 0, -1}

	var dist = make([][]int, h)
	for i := 0; i < h; i++ {
		dist[i] = make([]int, w)
		for j := 0; j < w; j++ {
			dist[i][j] = inf
		}
	}

	q := queue2{}
	dist[sx][sy] = 0
	q.push(point{sx, sy})

	for !q.empty() {
		u := q.pop()
		for i := 0; i < 4; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]
			if 0 <= tx && tx < h && 0 <= ty && ty < w && dist[tx][ty] == inf && dist[u.x][u.y]+1 <= dist[tx][ty] {
				dist[tx][ty] = dist[u.x][u.y] + 1
				q.push(point{tx, ty})
			}
		}
	}

	return dist
}

// 経路復元
func restore(h, w int, x, y int, dist [][]int) []point {
	r := []point{{x, y}}
	var dx = [4]int{1, 0, -1, 0}
	var dy = [4]int{0, 1, 0, -1}

	c := dist[x][y]

	for c > 0 {
		for i := 0; i < 4; i++ {
			tx := x + dx[i]
			ty := y + dy[i]

			if 0 <= tx && tx < h && 0 <= ty && ty < w && dist[tx][ty] == c-1 {
				r = append(r, point{tx, ty})
				c--
				x, y = tx, ty
			}
		}
	}

	return r
}
