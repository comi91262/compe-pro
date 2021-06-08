package main

import (
	"bufio"
	"fmt"
	"os"
)

var g [][]int

func dfs(n int, node *[]int) {
	if (*node)[n] > 0 {
		return
	}

	(*node)[n] = 1

	for _, b := range g[n] {
		dfs(b, node)
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
	for i := 1; i < n+1; i++ {
		var node = make([]int, n+1)
		dfs(i, &node)

		for j := 1; j < n+1; j++ {
			if node[j] > 0 {
				sum++
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", sum)
}
