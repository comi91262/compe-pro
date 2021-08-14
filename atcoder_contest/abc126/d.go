package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type edge struct {
	to   int
	cost int
}

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

var ans []int

func bfs(start int, g [][]edge) {
	q := queue{}
	q.push(start)

	used := make([]bool, len(g))
	used[start] = true
	for !q.empty() {
		node := q.pop()
		pre := ans[node]
		for _, e := range g[node] {
			ans[e.to] = pre ^ e.cost

			if !used[e.to] {
				q.push(e.to)
				used[node] = true
			}
		}
	}
}

func main() {
	defer writer.Flush()
	var n int

	fmt.Fscan(reader, &n)

	var g = make([][]edge, n+1)
	for i := 1; i < n+1; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		g[u] = append(g[u], edge{v, w % 2})
		g[v] = append(g[v], edge{u, w % 2})
	}
	ans = make([]int, n+1)

	bfs(1, g)
	for i := range ans {
		if i == 0 {
			continue
		}
		fmt.Fprintf(writer, "%v\n", ans[i])
	}
}
