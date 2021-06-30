package algo

func dfsTree(n, pre int, g *[][]int) {
	for _, next := range (*g)[n] {
		if pre == next {
			continue
		}

		dfs(next, n, g)
	}
}

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
		dfs(sx, sy, nx, ny, used)
	}
	used[x][y] = false
}
