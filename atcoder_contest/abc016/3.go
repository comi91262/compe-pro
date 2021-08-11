package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func contains(x int, a []int) bool {
	for _, e := range a {
		if x == e {
			return true
		}
	}
	return false
}

func dfs(n, pre, d int, f []int, ff *[]int, g [][]int) int {
	if d == 2 {
		if contains(n, f) || contains(n, *ff) {
			return 0
		} else {
			*ff = append(*ff, n)
			return 1
		}
	}

	r := 0
	for _, next := range g[n] {
		if next == pre {
			continue
		}

		r += dfs(next, n, d+1, f, ff, g)
	}

	return r
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var g = make([][]int, n+1)
	var a = make([]int, m+1)
	var b = make([]int, m+1)
	for i := 1; i < m+1; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}
	// fmt.Fprintf(writer, "%v\n", g)

	for i := 1; i <= n; i++ {
		f := g[i]
		f = append(f, i)
		ff := []int{}
		fmt.Fprintf(writer, "%v\n", dfs(i, i, 0, f, &ff, g))
	}
}
