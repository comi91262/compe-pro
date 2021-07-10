package algo

// DFSするだけ
func dfsTree(n, pre int, g *[][]int) {
	for _, next := range (*g)[n] {
		if pre == next {
			continue
		}

		dfsTree(next, n, g)
	}
}

// 1, 2 で二部グラフを彩色する関数
func dfsColor(n, pre int, color []int, g [][]int) {
	color[n] = pre
	for _, next := range g[n] {
		if color[next] > 0 {
			continue
		}
		dfsColor(next, 3-pre, color, g)
	}
}

// 二次元配列をDFSする バックトラック付き
func dfs2d(sx, sy, x, y int, used [][]bool) {
	var dx = [4]int{0, 1, 0, -1}
	var dy = [4]int{1, 0, -1, 0}

	if x == sx && y == sy && used[x][y] {
		return
	}

	h := len(used)
	w := len(used[0])

	used[x][y] = true

	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || ny < 0 || nx >= h || ny >= w {
			continue
		}
		if (nx != sx || ny != sy) && used[nx][ny] {
			continue
		}
		dfs2d(sx, sy, nx, ny, used)
	}
	used[x][y] = false
}
