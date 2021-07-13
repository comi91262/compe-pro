package algo

func getIdentityMatrix(n int) [][]int {
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
		c[i][i] = 1
	}

	return c
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
