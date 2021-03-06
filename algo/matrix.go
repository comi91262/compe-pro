package algo

func getIdentityMatrix(n int) [][]int {
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
		c[i][i] = 1
	}

	return c
}

// in-place  n * n
func rotateMatrix(a [][]string) {
	n := len(a)
	for x := 0; x < n/2; x++ {
		for y := x; y < n-x-1; y++ {
			a[x][y], a[y][n-1-x], a[n-1-x][n-1-y], a[n-1-y][x] = a[y][n-1-x], a[n-1-x][n-1-y], a[n-1-y][x], a[x][y]
		}
	}
}

func rotateClockwise(n, m int, g [][]string) (h [][]string) {
	h = make([][]string, m)
	for i := 0; i < m; i++ {
		h[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			h[j][n-i-1] = g[i][j]
		}
	}
	return
}

// 平行移動量を計算する
func diffMatrix(s [][]string, t [][]string) (int, int) {
	n := len(s)

	sx, sy := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == "#" {
				sx, sy = i, j
			}
		}

	}
	tx, ty := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if t[i][j] == "#" {
				tx, ty = i, j
			}
		}
	}

	return tx - sx, ty - sy
}

func prodMatrix(a, b [][]int) [][]int {
	aRowN := len(a)
	bColN := len(b[0])
	bRowN := len(b)

	c := make([][]int, aRowN)
	for i := 0; i < aRowN; i++ {
		c[i] = make([]int, bColN)
	}

	for i := 0; i < aRowN; i++ {
		for k := 0; k < bRowN; k++ {
			for j := 0; j < bColN; j++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return c
}

func prodModMatrix(a, b [][]int, mod int) [][]int {
	aRowN := len(a)
	bColN := len(b[0])
	bRowN := len(b)

	c := make([][]int, aRowN)
	for i := 0; i < aRowN; i++ {
		c[i] = make([]int, bColN)
	}

	for i := 0; i < aRowN; i++ {
		for k := 0; k < bRowN; k++ {
			for j := 0; j < bColN; j++ {
				c[i][j] = (c[i][j] + a[i][k]*b[k][j]) % mod
			}
		}
	}

	return c
}

func powMatrix(a [][]int, x int) [][]int {
	n := len(a)

	r := getIdentityMatrix(n)

	for x > 0 {
		if x&1 == 1 {
			r = prodMatrix(r, a)
		}
		a = prodMatrix(a, a)
		x >>= 1
	}
	return r
}

func powModMatrix(a [][]int, x int, mod int) [][]int {
	n := len(a)

	r := getIdentityMatrix(n)

	for x > 0 {
		if x&1 == 1 {
			r = prodModMatrix(r, a, mod)
		}
		a = prodModMatrix(a, a, mod)
		x >>= 1
	}
	return r
}

// 2進数の掃き出し法, 破壊的
// 1 0 1 1 0  ->  1 0 0 1 1
// 1 0 0 1 1      0 1 0 1 1
// 0 1 1 1 0      0 0 1 0 1
// 0 0 1 0 1      0 0 0 0 0
// compression(compression.go) が必要
func sweepBitMatrix(a [][]int) [][]int {
	if row := len(a); row == 0 {
		return nil
	}

	row := len(a)
	col := len(a[0])

	cnt := make([]int, row)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if a[i][j] == 0 {
				cnt[i]++
				continue
			}
			break
		}

		if cnt[i] == col {
			continue
		}

		for j := 0; j < row; j++ {
			if i == j || a[j][cnt[i]] == 0 {
				continue
			}

			for k := 0; k < col; k++ {
				a[j][k] ^= a[i][k]
			}
		}
	}

	cnt = compression(cnt)

	c := make([][]int, row)
	for i := 0; i < row; i++ {
		c[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		c[cnt[i]] = a[i]
	}

	return c
}

// 上下反転
func reverseMatrix(a [][]int) [][]int {
	n := len(a)
	m := len(a[0])

	for i := 0; i < n/2; i++ {
		for j := 0; j < m; j++ {
			a[i][j], a[n-1-i][j] = a[n-1-i][j], a[i][j]
		}
	}

	return a
}
