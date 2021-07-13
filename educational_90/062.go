package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type queue []int

func (q *queue) push(n int) {
	*q = append(*q, n)
}

func (q *queue) pop() int {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n+1)
	var b = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	g := make([][]int, n+1)
	used := make([]bool, n+1)
	q := queue{}
	for i := 1; i <= n; i++ {
		x, y := a[i], b[i]
		g[x] = append(g[x], i)
		g[y] = append(g[y], i)

		if y == i || x == i {
			q.push(i)
			used[i] = true
		}
	}

	ans := []int{}
	for !q.empty() {
		node := q.pop()
		ans = append(ans, node)
		for _, e := range g[node] {
			if !used[e] {
				used[e] = true
				q.push(e)
			}

		}
	}

	if len(ans) != n {
		fmt.Fprintf(writer, "%v\n", -1)
		return
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", ans[n-1-i])

	}
}
