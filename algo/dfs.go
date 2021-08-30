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

// 重複した数列を作る
func dfsD(n, m, d, k int, a []int, accum *[][]int) {
	if n == d {
		b := make([]int, len(a))
		copy(b, a)
		*accum = append(*accum, b)
		return
	}

	for i := 0; k+i <= m; i++ {
		a = append(a, k+i)
		dfsD(n, m, d+1, k+i, a, accum)
		a = a[:len(a)-1]
	}
}

// 二次元配列をDFSする バックトラック付き
func dfs2d(x, y int, s [][]string, used [][]bool) {
	if used[x][y] {
		return
	}

	var dx = [4]int{0, 1, 0, -1}
	var dy = [4]int{1, 0, -1, 0}

	h := len(used)
	w := len(used[0])

	used[x][y] = true

	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || ny < 0 || nx >= h || ny >= w {
			continue
		}
		dfs2d(x, y, s, used)
	}
	used[x][y] = false
}
