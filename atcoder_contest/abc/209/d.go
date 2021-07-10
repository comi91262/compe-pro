package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func dfs(n, pre int, color []int, g [][]int) {
	color[n] = pre
	for _, next := range g[n] {
		if color[next] > 0 {
			continue
		}
		dfs(next, 3-pre, color, g)
	}
}

func main() {
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	var g = make([][]int, n+1)
	var a = make([]int, n)
	var b = make([]int, n)
	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	var c = make([]int, q)
	var d = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &c[i], &d[i])
	}

	var color = make([]int, n+1)
	dfs(1, 1, color, g)

	for i := 0; i < q; i++ {
		if color[c[i]] == color[d[i]] {
			fmt.Fprintf(writer, "%v\n", "Town")
		} else {
			fmt.Fprintf(writer, "%v\n", "Road")
		}
	}

}
