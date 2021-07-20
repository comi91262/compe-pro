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

func (d *Dinic) addEdge(from, to, cap int) {
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

func (d *Dinic) maxFlow(s, t int) int {
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

type pair struct {
	first  int
	second int
}

var dx = [8]int{1, 1, 0, -1, -1, -1, 0, 1}
var dy = [8]int{0, 1, 1, 1, 0, -1, -1, -1}

func main() {
	defer writer.Flush()

	// Step #1. Input
	var n, t int
	fmt.Fscan(reader, &n, &t)

	var nex [1 << 18][8]int
	var ans [1 << 18]int

	var ax = make([]int, n+1)
	var ay = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &ax[i], &ay[i])
	}

	var bx = make([]int, n+1)
	var by = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &bx[i], &by[i])
	}

	mapper := make(map[pair]int)
	for i := 1; i <= n; i++ {
		mapper[pair{first: bx[i], second: by[i]}] = i
	}
	fmt.Fprintf(writer, "%v\n", n)
	fmt.Fprintf(writer, "%v\n", t)
	fmt.Fprintf(writer, "%v %v\n", ax, ay)
	fmt.Fprintf(writer, "%v %v\n", bx, by)
	fmt.Fprintf(writer, "%v\n", mapper)

	d := Dinic{
		g:     make([][]edge, 2*n+5),
		level: make([]int, 2*n+5),
		iter:  make([]int, 2*n+5),
	}

	// Step #2. Add edge
	for i := 1; i <= n; i++ {
		for j := 0; j < 8; j++ {
			tx := ax[i] + dx[j]*t
			ty := ay[i] + dy[j]*t
			nex[i][j] = mapper[pair{first: tx, second: ty}]

			if nex[i][j] != 0 {
				fmt.Fprintf(writer, "%v %v %v\n", tx, ty, mapper[pair{first: tx, second: ty}])
				d.addEdge(i, n+nex[i][j], 1)
			}
		}
	}
	for i := 1; i < +n; i++ {
		d.addEdge(2*n+1, i, 1)
	}
	for i := 1; i < +n; i++ {
		d.addEdge(i+n, 2*n+2, 1)
	}

	// Step #3. Max Flow
	res := d.maxFlow(2*n+1, 2*n+2)
	res = d.maxFlow(2*n+1, 2*n+2)
	res = d.maxFlow(2*n+1, 2*n+2)
	fmt.Fprintf(writer, "%v\n", res)
	// if res != n {
	// 	fmt.Fprintf(writer, "%v\n", "No")
	// 	return
	// }

	// Step #4. get ans
	for i := 1; i <= n; i++ {
		for j := 0; j < len(d.g[i]); j++ {
			if d.g[i][j].to > 2*n || d.g[i][j].cap == 1 {
				continue
			}
			for k := 0; k < 8; k++ {
				if nex[i][k] == d.g[i][j].to-n {
					ans[i] = k + 1
				}
			}
		}
	}

	// Step #5. Output
	fmt.Fprintf(writer, "%v\n", "Yes")
	for i := 1; i <= n; i++ {
		if i >= 2 {
			fmt.Fprintf(writer, " ")
		}
		fmt.Fprintf(writer, "%v", ans[i])
	}
	fmt.Fprintf(writer, "\n")
}
