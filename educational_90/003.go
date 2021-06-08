package main

import (
	"bufio"
	"fmt"
	"os"
)

var g [][]int
var memo []int

func dfs(n int, d int) (int, int) {
	memo[n]++

	mn := n
	md := d
	for _, next := range g[n] {
		if memo[next] > 0 {
			continue
		}

		var md2, mn2 = dfs(next, d+1)
		if md < md2 {
			md = md2
			mn = mn2
		}
	}

	return md, mn
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	g = make([][]int, n+1)

	var x, y int
	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &x, &y)
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	memo = make([]int, n+1)
	var _, mn = dfs(1, 0)

	memo = make([]int, n+1)
	var md, _ = dfs(mn, 0)

	fmt.Fprintf(writer, "%v\n", md+1)
}
