package algo

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
		update := 1
		for update > 0 {
			update = 0
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
						update = 1
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
