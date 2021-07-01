package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var used []bool

func dfs1(x int, g [][]int, path *[]int) {
	used[x] = true

	for _, next := range g[x] {
		if used[next] {
			continue
		}

		dfs1(next, g, path)
	}

	*path = append(*path, x)
}

func dfs2(x int, g [][]int) int {
	used[x] = true
	r := 0
	for _, next := range g[x] {
		if used[next] {
			continue
		}

		r += dfs2(next, g)
	}

	return r + 1
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	var g = make([][]int, n+1)
	var h = make([][]int, n+1)

	var a = make([]int, m+1)
	var b = make([]int, m+1)
	for i := 1; i < m+1; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		h[b[i]] = append(h[b[i]], a[i])
	}

	path := make([]int, 0, n)
	used = make([]bool, n+1)
	for i := 1; i < n+1; i++ {
		if used[i] {
			continue
		}
		dfs1(i, g, &path)
	}

	ans := 0
	used = make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if used[path[n-i]] {
			continue
		}
		group := make([]int, 0)

		dfs1(path[n-i], h, &group)
		ans += len(group) * (len(group) - 1) / 2
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
