package maxflow

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

type pair struct {
	first  int
	second int
}

type Edge struct {
	from int
	to   int
	cap  int
	flow int
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
	pos   []pair
}

func (d *Dinic) AddEdge(from, to, cap int) {
	d.pos = append(d.pos, pair{from, len(d.g[from])})
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
		for i := 0; i < len(d.g[v]); i++ {
			e := &d.g[v][i]
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

func (d *Dinic) getEdge(i int) Edge {
	e := d.g[d.pos[i].first][d.pos[i].second]
	re := d.g[e.to][e.rev]
	return Edge{d.pos[i].first, e.to, e.cap + re.cap, re.cap}
}

func (d *Dinic) Edges() []Edge {
	m := len(d.pos)
	result := []Edge{}
	for i := 0; i < m; i++ {
		result = append(result, d.getEdge(i))
	}
	return result
}

func (d *Dinic) Create(n int) {
	d.g = make([][]edge, n)
	d.level = make([]int, n)
	d.iter = make([]int, n)
}
