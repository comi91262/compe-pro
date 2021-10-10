package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func scanInt() int       { sc.Scan(); return parseInt(sc.Bytes()) }
func scanString() string { sc.Scan(); return string(sc.Bytes()) }
func scanFloat() float64 { sc.Scan(); return parseFloat(sc.Bytes()) }
func scanInts(n int) (ints []int) {
	ints = make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = scanInt()
	}
	return
}
func scanStrings(n int) (st []string) {
	st = make([]string, n)
	for i := 0; i < n; i++ {
		st[i] = scanString()
	}
	return
}
func scanFloats(n int) (fs []float64) {
	fs = make([]float64, n)
	for i := 0; i < n; i++ {
		fs[i] = scanFloat()
	}
	return fs
}

func parseInt(b []byte) (n int) {
	for _, ch := range b {
		ch -= '0'
		if ch <= 9 {
			n = n*10 + int(ch)
		}
	}
	if b[0] == '-' {
		return -n
	}
	return
}

var float64pow10 = []float64{
	1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6,
	1e7, 1e8, 1e9, 1e10, 1e11, 1e12,
	1e13, 1e14, 1e15, 1e16, 1e17, 1e18,
	1e19, 1e20, 1e21, 1e22,
}

func parseFloat(b []byte) float64 {
	f, dot := 0.0, 0
	for i, ch := range b {
		if ch == '.' {
			dot = i + 1
			continue
		}
		if ch -= '0'; ch <= 9 {
			f = f*10 + float64(ch)
		}
	}
	if dot != 0 {
		f /= float64pow10[len(b)-dot]
	}
	if b[0] == '-' {
		return -f
	}
	return f
}

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
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
func larger(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
func smaller(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)

	n, m := scanInt(), scanInt()
	grid := make([][]string, n)
	for i := 0; i < n; i++ {
		grid[i] = strings.Split(scanString(), "")
	}

	d := Dinic{}
	d.Create(n*m + 2)

	s := n * m
	t := n*m + 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == "#" {
				continue
			}
			v := i*m + j
			if (i+j)%2 == 0 {
				d.AddEdge(s, v, 1)
			} else {
				d.AddEdge(v, t, 1)
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == "#" || (i+j)%2 == 1 {
				continue
			}
			v := i*m + j
			if i > 0 {
				if grid[i-1][j] == "." {
					d.AddEdge(v, (i-1)*m+j, 1)
				}
			}
			if j > 0 {
				if grid[i][j-1] == "." {
					d.AddEdge(v, i*m+j-1, 1)
				}
			}
			if i < n-1 {
				if grid[i+1][j] == "." {
					d.AddEdge(v, (i+1)*m+j, 1)
				}
			}
			if j < m-1 {
				if grid[i][j+1] == "." {
					d.AddEdge(v, i*m+j+1, 1)
				}
			}
		}
	}

	fmt.Fprintf(wr, "%v\n", d.MaxFlow(n*m, n*m+1))
	ans := make([][]string, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]string, m)
		for j := 0; j < m; j++ {
			ans[i][j] = grid[i][j]
		}
	}

	for _, e := range d.Edges() {
		if e.from == s || e.to == t || e.flow == 0 {
			continue
		}
		fi, fj := e.from/m, e.from%m
		ti, tj := e.to/m, e.to%m
		if fi == ti {
			ans[fi][smaller(fj, tj)] = ">"
			ans[fi][larger(fj, tj)] = "<"
		} else {
			ans[smaller(fi, ti)][fj] = "v"
			ans[larger(fi, ti)][fj] = "^"
		}
	}

	for i := range ans {
		fmt.Fprintf(wr, "%v\n", strings.Join(ans[i], ""))
	}

}
