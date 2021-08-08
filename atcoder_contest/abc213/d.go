package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func dfs(n, pre int, used []bool, g [][]int) {
	used[n] = true
	fmt.Fprintf(writer, "%v\n", n)

	sort.Ints(g[n])
	for _, next := range g[n] {
		if used[next] {
			continue
		}

		dfs(next, n, used, g)
		fmt.Fprintf(writer, "%v\n", n)
	}
	// if len(g[n]) != 1 || g[n][0] != pre {
	//fmt.Fprintf(writer, "%v\n", n)
	// }
	used[n] = false
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var g = make([][]int, n+1)
	var a = make([]int, n)
	var b = make([]int, n)
	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	var used = make([]bool, n+1)
	dfs(1, 1, used, g)
}
