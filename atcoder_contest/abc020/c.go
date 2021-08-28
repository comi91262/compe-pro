package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

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

func bfs(h, w int, sx, sy, cost int, s [][]string) [][]int {
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

	q := create(h * w * 4)

	dist[sx][sy] = 0
	q.pushFront(point{sx, sy})

	for !q.empty() {
		u := q.popFront()
		for i := 0; i < 4; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]
			if 0 <= tx && tx < h && 0 <= ty && ty < w {
				if s[tx][ty] == "#" {
					d := dist[u.x][u.y] + cost
					if d < dist[tx][ty] {
						dist[tx][ty] = d
						q.pushBack(point{tx, ty})
					}
				} else {
					d := dist[u.x][u.y] + 1
					if d < dist[tx][ty] {
						dist[tx][ty] = d
						q.pushFront(point{tx, ty})
					}
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
	var t int
	fmt.Fscan(reader, &t)

	sx, sy := 0, 0
	tx, ty := 0, 0

	s := make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == "S" {
				sx, sy = i, j
				s[i][j] = "."
				continue
			}

			if s[i][j] == "G" {
				tx, ty = i, j
				s[i][j] = "."
				continue
			}
		}
	}

	ng := 1_000_000_000
	ok := -1

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		dist := bfs(h, w, sx, sy, mid, s)
		if dist[tx][ty] <= t {
			ok = mid
		} else {
			ng = mid
		}
	}

	fmt.Fprintf(writer, "%v\n", ok)
}
