package mincostflow

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
