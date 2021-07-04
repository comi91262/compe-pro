package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const inf = 1_000_000_000 + 7

var dist [1009][1009][4]int
var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}

type state struct {
	x   int
	y   int
	dir int
}

func main() {
	defer writer.Flush()

	var h, w int
	fmt.Fscan(reader, &h, &w)

	var rs, cs, rt, ct int
	fmt.Fscan(reader, &rs, &cs, &rt, &ct)
	rs--
	cs--
	rt--
	ct--

	var s = make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			dist[i][j][0] = inf
			dist[i][j][1] = inf
			dist[i][j][2] = inf
			dist[i][j][3] = inf
		}
	}

	deq := create(4080000)
	for i := 0; i < 4; i++ {
		dist[rs][cs][i] = 0
		deq.pushFront(state{rs, cs, i})
	}

	for !deq.empty() {
		u := deq.popFront()
		for i := 0; i < 4; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]
			cost := dist[u.x][u.y][u.dir]
			if u.dir != i {
				cost += 1
			}
			if 0 <= tx && tx < h && 0 <= ty && ty < w && s[tx][ty] == "." && dist[tx][ty][i] > cost {
				dist[tx][ty][i] = cost
				if u.dir != i {
					deq.pushBack(state{tx, ty, i})
				} else {
					deq.pushFront(state{tx, ty, i})
				}
			}
		}
	}
	ans := inf
	for i := 0; i < 4; i++ {
		ans = min(ans, dist[rt][ct][i])
	}
	fmt.Fprintf(writer, "%v\n", ans)
}

// var n, q int
// fmt.Fscan(reader, &n, &q)

// var s string
// fmt.Fscan(reader, &n, &k, &s)
// ss := strings.Split(s, "")

// var a = make([]int, n)
// for i := 0; i < n; i++ {
// 	fmt.Fscan(reader, &a[i])
// }
// var n int
// fmt.Fscan(reader, &n)
// fmt.Fprintf(writer, "%d\n", n)

//	var n int
//	fmt.Fscan(reader, &n)
//	var a = make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Fscan(reader, &a[i])
//	}
// const d = 1_000_000_000 + 7
// for i := 0; i < n; i++ {
// for i := 0; i < n; i++ {
// }
// }

// var g = make([][]int, n+1)
// var a = make([]int, m+1)
// var b = make([]int, m+1)
// for i := 1; i < m+1; i++ {
// 	fmt.Fscan(reader, &a[i], &b[i])
// 	g[a[i]] = append(g[a[i]], b[i])
// 	g[b[i]] = append(g[b[i]], a[i])
// }

// m := map[int]int{}
// a := make([]int, n)
// for i := 0; i < len(a); i++ {
// 	m[a[i]]++
// }

// func abs(x int) int {
// 	if x >= 0 {
// 		return x
// 	} else {
// 		return x * -1
// 	}
// }
//
func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

//
// func max(arg ...int) int {
// 	max := arg[0]
// 	for _, x := range arg {
// 		if max < x {
// 			max = x
// 		}
// 	}
// 	return max
// }

type Deque struct {
	cl   int
	cr   int
	data []state
}

func (d *Deque) pushFront(x state) {
	d.cl--
	d.data[d.cl] = x
}

func (d *Deque) pushBack(x state) {
	d.data[d.cr] = x
	d.cr++
}

func (d *Deque) popFront() state {
	v := d.data[d.cl]
	d.cl++
	return v
}

func (d *Deque) popBack() state {
	v := d.data[d.cr-1]
	d.cr--
	return v
}

func (d *Deque) size() int {
	return d.cr - d.cl
}

func (d *Deque) empty() bool {
	return d.size() == 0
}

func (d *Deque) get(x int) state {
	return d.data[d.cl+x-1]
}

func create(size int) Deque {
	return Deque{cl: size, cr: size, data: make([]state, size*2, size*2)}
}
