package algo

type point struct {
	x int
	y int
}

type queue []point

func (q *queue) push(n point) {
	*q = append(*q, n)
}

func (q *queue) pop() point {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) front() point {
	return (*q)[0]
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

// 単一コストの最短経路問題を BFS でとく (二次元版)
func bfs2(h, w, sx, sy int, s [][]string) [][]int {
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

	q := queue{}
	dist[sx][sy] = 0
	q.push(point{sx, sy})

	for !q.empty() {
		u := q.pop()
		for i := 0; i < 4; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]

			if tx < 0 || tx >= h || ty < 0 || ty >= w {
				continue
			}
			if dist[tx][ty] != inf || s[tx][ty] == "#" {
				continue
			}
			if dist[u.x][u.y]+1 <= dist[tx][ty] {
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
