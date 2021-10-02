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

// x (lower <= x <= upper) を使った数列を全列挙する。重複あり
// 0 <= x <= 2, size = 2 -> 00,01,02,10,11,12,20,21,22
func makeDuplicatedNumberArray(size, depth, lower, upper int, a []int, accum *[][]int) {
	if size == depth {
		b := make([]int, len(a))
		copy(b, a)
		*accum = append(*accum, b)
		return
	}

	for i := lower; i <= upper; i++ {
		a = append(a, i)
		makeDuplicatedNumberArray(size, depth+1, lower, upper, a, accum)
		a = a[:len(a)-1]
	}
}

// x (lower <= x <= upper) を使った数列を全列挙する。重複なし
// 0 <= x <= 2, size = 2 -> 01,02,10,12,20,21
func makeNumberArray(size, depth, lower, upper int, memo map[int]int, a []int, accum *[][]int) {
	if size == depth {
		b := make([]int, len(a))
		copy(b, a)
		*accum = append(*accum, b)
		return
	}

	for i := lower; i <= upper; i++ {
		if memo[i] > 0 {
			continue
		}
		a = append(a, i)
		memo[i]++
		makeNumberArray(size, depth+1, lower, upper, memo, a, accum)
		a = a[:len(a)-1]
		memo[i]--
	}
}

func dfsDst(n, d int, a string, accum *[]string) {
	if 0 < d && d <= n {
		*accum = append(*accum, a)
	}
	if d == n {
		return
	}

	for i := 0; i < 26; i++ {
		a += string("a"[0] + byte(i))
		dfsDst(n, d+1, a, accum)
		a = a[:len(a)-1]
	}
}

// 二次元配列をDFSする
func dfs2d(h, w, x, y int, s [][]string, used [][]bool) {
	if used[x][y] {
		return
	}

	var dx = [4]int{0, 1, 0, -1}
	var dy = [4]int{1, 0, -1, 0}

	used[x][y] = true
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || ny < 0 || nx >= h || ny >= w {
			continue
		}
		if s[nx][ny] != "#" {
			continue
		}
		dfs2d(h, w, nx, ny, s, used)
	}
	used[x][y] = false
}
