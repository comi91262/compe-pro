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
	var q int
	fmt.Fscan(reader, &q)

	var g = make([][]int, n)
	var a = make([]int, m)
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		a[i]--
		b[i]--
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	var c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	for i := 0; i < q; i++ {
		var t, x, y int
		fmt.Fscan(reader, &t, x, y)
		switch t {
		case 1:
			fmt.Fscan(reader, &x)
			x--
			fmt.Fprintf(writer, "%v\n", c[x])
			for j := range g[x] {
				c[g[x][j]] = c[x]
			}
		case 2:
			fmt.Fscan(reader, &x, &y)
			x--
			fmt.Fprintf(writer, "%v\n", c[x])
			c[x] = y
		}
	}
}
