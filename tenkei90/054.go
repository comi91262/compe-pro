package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const inf = -2

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

	var n, m int
	fmt.Fscan(reader, &n, &m)

	r := make([][]int, m)
	for i := 0; i < m; i++ {
		var k int
		fmt.Fscan(reader, &k)

		r[i] = []int{}
		for j := 0; j < k; j++ {
			var tmp int
			fmt.Fscan(reader, &tmp)
			r[i] = append(r[i], tmp)
		}
	}

	g := make([][]int, n+1)
	node := n + 1
	for i := 0; i < m; i++ {
		l := len(r[i])
		tmp := []int{}
		for j := 0; j < l; j++ {
			tmp = append(tmp, r[i][j])
		}

		g = append(g, tmp)
		for _, e := range tmp {
			g[e] = append(g[e], node)
		}
		node++
	}

	q := queue{}
	q.push(1)

	dist := make([]int, node)

	for i := 0; i < node; i++ {
		dist[i] = inf
	}
	dist[1] = 0

	for !q.empty() {
		node := q.pop()
		for _, e := range g[node] {
			if dist[e] == inf {
				dist[e] = dist[node] + 1
				q.push(e)
			}
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintf(writer, "%v\n", dist[i]/2)
	}
}
