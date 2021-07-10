package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func getIdentityMatrix(n int) [][]int {
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
		c[i][i] = 1
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

func main() {
	defer writer.Flush()

	var p, q, r, k int
	fmt.Fscan(reader, &p, &q, &r, &k)

	a := [][]int{
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
	}

	b := powModMatrix(a, k-3, 10)
	p %= 10
	q %= 10
	r %= 10

	c := prodModMatrix(b, [][]int{{r}, {q}, {p}}, 10)

	fmt.Fprintf(writer, "%v", c[0][0])
}
