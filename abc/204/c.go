package main

import (
	"bufio"
	"fmt"
	"os"
)

var dist [2001]struct{ ns []int }

func search(n int, node *[]int) {
	var slice = dist[n].ns

	if (*node)[n] > 0 {
		return
	}

	(*node)[n] = 1

	for _, b := range slice {
		search(b, node)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a, &b)
		dist[a].ns = append(dist[a].ns, b)
	}

	sum := 0
	for i := 1; i < n+1; i++ {
		var node = make([]int, n+1)
		search(i, &node)

		for j := 1; j < n+1; j++ {
			if node[j] > 0 {
				sum++
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", sum)
}
