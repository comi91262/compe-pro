package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	var c int
	fmt.Fscan(reader, &c)

	var b = make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = make([]int, 1)
		for j := 0; j < 1; j++ {
			fmt.Fscan(reader, &b[i][j])
		}
	}
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	p := prodMatrix(a, b)
	ans := 0
	for i := 0; i < n; i++ {
		if p[i][0]+c > 0 {
			ans++
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
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
