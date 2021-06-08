package main

import (
	"fmt"
)

var g [][]int
var memo []bool

func dfs(n int) {
	if memo[n] {
		return
	}
	memo = true

	for _, next := range g[n] {
		dfs(next)
	}
}

func main() {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	g = make([][]int, n+1)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a, &b)
		g[a] = append(g[a], b)
	}

}
