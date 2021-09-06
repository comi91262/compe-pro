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

type edge struct {
	to   int
	cap  int
	rev  int
	cost int
}

type PrimalDual struct {
	n int
	g [][]edge
}

func (p *PrimalDual) Create(v int) {
	p.n = v
	p.g = make([][]edge, v)
}

func (p *PrimalDual) AddEdge(from, to, cap, cost int) {
	p.g[from] = append(p.g[from], edge{to: to, cap: cap, rev: len(p.g[to]), cost: cost})
	p.g[to] = append(p.g[to], edge{to: from, cap: 0, rev: len(p.g[from]) - 1, cost: -cost})
}

func (p *PrimalDual) MinCostFlow(s, t, f int) int {
	n := p.n
	prevv := make([]int, n)
	preve := make([]int, n)
	inf := 1 << 60
	r := 0
	for f > 0 {
		dist := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = inf
		}
		dist[s] = 0
		update := true
		for update {
			update = false
			for v := range make([]int, n) {
				if dist[v] == inf {
					continue
				}
				gv := p.g[v]
				for i := range make([]int, len(gv)) {
					to, cap, cost := gv[i].to, gv[i].cap, gv[i].cost
					if cap > 0 && dist[v]+cost < dist[to] {
						dist[to] = dist[v] + cost
						prevv[to] = v
						preve[to] = i
						update = true
					}
				}
			}
		}

		if dist[t] == inf {
			return -1
		}

		d := f
		for v := t; v != s; v = prevv[v] {
			d = min(d, p.g[prevv[v]][preve[v]].cap)
		}
		f -= d
		r += d * dist[t]
		for v := t; v != s; v = prevv[v] {
			p.g[prevv[v]][preve[v]].cap -= d
			p.g[v][p.g[prevv[v]][preve[v]].rev].cap += d
		}
	}
	return r
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	var q int
	fmt.Fscan(reader, &q)

	var w = make([]int, n)
	var v = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &w[i])
		fmt.Fscan(reader, &v[i])
	}

	var x = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &x[i])
	}
	var l = make([]int, q)
	var r = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &l[i])
		fmt.Fscan(reader, &r[i])
	}

	mx := 100000
	for i := 0; i < q; i++ {
		p := PrimalDual{}
		p.Create(n + m + 2)

		for j := 1; j <= n; j++ {
			p.AddEdge(0, j, 1, 0)
		}
		for j := n + 1; j <= n+m; j++ {
			p.AddEdge(j, n+m+1, 1, 0)
		}
		for j := 1; j <= n; j++ {
			for k := 1; k <= m; k++ {
				if l[i] <= k && k <= r[i] {
					continue
				}
				if w[j-1] <= x[k-1] {
					p.AddEdge(j, k+n, 1, mx-v[j-1])
				} else {
					p.AddEdge(j, k+n, 0, mx-v[j-1])
				}
			}
		}

		cur := 0
		for j := 1; j <= n; j++ {
			cost := p.MinCostFlow(0, n+m+1, 1)
			if cost == -1 {
				break
			}
			cur += mx - cost
		}
		fmt.Fprintf(writer, "%v\n", cur)
	}
}
