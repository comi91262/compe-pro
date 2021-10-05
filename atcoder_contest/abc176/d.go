package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type point struct {
	x int
	y int
}

type Deque struct {
	cl   int
	cr   int
	data []point
}

func (d *Deque) pushFront(x point) {
	d.cl--
	d.data[d.cl] = x
}

func (d *Deque) pushBack(x point) {
	d.data[d.cr] = x
	d.cr++
}

func (d *Deque) popFront() point {
	v := d.data[d.cl]
	d.cl++
	return v
}

func (d *Deque) popBack() point {
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

func (d *Deque) get(x int) point {
	return d.data[d.cl+x-1]
}

func create(size int) Deque {
	return Deque{cl: size, cr: size, data: make([]point, size*2, size*2)}
}

func bfs2MinDist(h, w, sx, sy int, s [][]string) [][]int {
	const inf = 1 << 60
	var dx = [4]int{1, 0, -1, 0}
	var dy = [4]int{0, 1, 0, -1}

	var dist = make([][]int, h)
	for i := 0; i < h; i++ {
		dist[i] = make([]int, w)
		for j := 0; j < w; j++ {
			dist[i][j] = inf
		}
	}

	q := create(h * w)
	dist[sx][sy] = 0
	q.pushFront(point{sx, sy})

	for !q.empty() {
		u := q.popFront()
		for i := 0; i < 4; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]

			if tx < 0 || tx >= h || ty < 0 || ty >= w {
				continue
			}
			if dist[tx][ty] == dist[u.x][u.y] || s[tx][ty] == "#" {
				continue
			}
			if dist[u.x][u.y] <= dist[tx][ty] {
				dist[tx][ty] = dist[u.x][u.y]
				q.pushFront(point{tx, ty})
			}
		}
		for i := -2; i <= 2; i++ {
			for j := -2; j <= 2; j++ {
				tx := u.x + i
				ty := u.y + j
				if tx < 0 || tx >= h || ty < 0 || ty >= w {
					continue
				}
				if dist[tx][ty] != inf || s[tx][ty] == "#" {
					continue
				}
				if dist[u.x][u.y]+1 <= dist[tx][ty] {
					dist[tx][ty] = dist[u.x][u.y] + 1
					q.pushBack(point{tx, ty})
				}
			}
		}

	}

	return dist
}

func main() {
	defer writer.Flush()

	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	var ch int
	fmt.Fscan(reader, &ch)
	var cw int
	fmt.Fscan(reader, &cw)
	var dh int
	fmt.Fscan(reader, &dh)
	var dw int
	fmt.Fscan(reader, &dw)
	ch--
	cw--
	dh--
	dw--

	s := make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	dist := bfs2MinDist(h, w, ch, cw, s)
	if dist[dh][dw] == 1<<60 {
		fmt.Fprintf(writer, "%v\n", -1)
	} else {
		fmt.Fprintf(writer, "%v\n", dist[dh][dw])
	}
}
func printer(a [][]int, limit int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] >= limit {
				fmt.Fprintf(writer, "%5v ", "x")
				continue
			}
			fmt.Fprintf(writer, "%5v ", a[i][j])
		}
		fmt.Fprintf(writer, "\n")
	}
}
