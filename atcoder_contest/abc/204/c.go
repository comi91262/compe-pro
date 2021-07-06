package main

import (
	"bufio"
	"fmt"
	"os"
)

var g [][]int
var memo []bool

func dfs(n int) {
	if memo[n] {
		return
	}

	memo[n] = true

	for _, next := range g[n] {
		dfs(next)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	g = make([][]int, n+1)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a, &b)
		g[a] = append(g[a], b)
	}

	sum := 0
	memo = make([]bool, n+1)
	for i := 1; i < n+1; i++ {
		for j := 0; j < len(memo); j++ {
			memo[j] = false
		}

		dfs(i)

		for j := 1; j < n+1; j++ {
			if memo[j] {
				sum++
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", sum)
}
