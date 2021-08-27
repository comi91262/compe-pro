package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
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

type edge struct {
	to  int
	cap int
	rev int
}

type Dinic struct {
	g     [][]edge
	level []int
	iter  []int
}

func (d *Dinic) AddEdge(from, to, cap int) {
	d.g[from] = append(d.g[from], edge{to: to, cap: cap, rev: len(d.g[to])})
	d.g[to] = append(d.g[to], edge{to: from, cap: 0, rev: len(d.g[from]) - 1})

}

func (d *Dinic) bfs(s int) {
	for i := range d.level {
		d.level[i] = -1
	}

	q := queue{}
	d.level[s] = 0

	q.push(s)
	for !q.empty() {
		v := q.pop()
		for _, e := range d.g[v] {
			if e.cap > 0 && d.level[e.to] < 0 {
				d.level[e.to] = d.level[v] + 1
				q.push(e.to)
			}
		}
	}
}

func (d *Dinic) dfs(v, t, f int) int {
	if v == t {
		return f
	}
	for i := d.iter[v]; i < len(d.g[v]); i++ {
		e := &d.g[v][i]
		if e.cap > 0 && d.level[v] < d.level[e.to] {
			f := d.dfs(e.to, t, min(f, e.cap))
			if f > 0 {
				e.cap -= f
				d.g[e.to][e.rev].cap += f
				return f
			}
		}
		d.iter[v]++
	}

	return 0
}

func (d *Dinic) MaxFlow(s, t int) int {
	flow := 0

	for {
		d.bfs(s)
		if d.level[t] < 0 {
			return flow
		}
		for i := range d.iter {
			d.iter[i] = 0
		}
		for {
			f := d.dfs(s, t, 1<<60)
			if f <= 0 {
				break
			}
			flow += f
		}
	}
}

func (d *Dinic) Create(n int) {
	d.g = make([][]edge, 2*n+2)
	d.level = make([]int, 2*n+2)
	d.iter = make([]int, 2*n+2)
}

type pair struct {
	first  int
	second int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var x = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i].first, &x[i].second)
	}

	var y = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &y[i].first, &y[i].second)
	}

	d := Dinic{
		g:     make([][]edge, 2*n+2),
		level: make([]int, 2*n+2),
		iter:  make([]int, 2*n+2),
	}
	for i := 1; i <= n; i++ {
		d.AddEdge(0, i, 1)
	}
	for i := n + 1; i <= n+n; i++ {
		d.AddEdge(i, 2*n+1, 1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if x[i].first < y[j].first && x[i].second < y[j].second {
				d.AddEdge(i+1, j+n+1, 1)
			}
		}
	}

	fmt.Fprintf(writer, "%v\n", d.MaxFlow(0, 2*n+1))

}
