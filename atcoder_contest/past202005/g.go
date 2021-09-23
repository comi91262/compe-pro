package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type point struct {
	x int
	y int
}
type queue []point

func (q *queue) push(n point) {
	*q = append(*q, n)
}

func (q *queue) pop() point {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func bfs2MinDist(h, w, sx, sy int, s [][]string) [][]int {
	const inf = 1 << 60
	var dx = [6]int{1, 0, -1, 0, -1, -1}
	var dy = [6]int{0, 1, 0, -1, 1, -1}

	var dist = make([][]int, h)
	for i := 0; i < h; i++ {
		dist[i] = make([]int, w)
		for j := 0; j < w; j++ {
			dist[i][j] = inf
		}
	}

	q := queue{}
	dist[sx][sy] = 0
	q.push(point{sx, sy})

	for !q.empty() {
		u := q.pop()
		for i := 0; i < 6; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]

			if tx < 0 || tx >= h || ty < 0 || ty >= w {
				continue
			}
			if dist[tx][ty] != inf || s[tx][ty] == "#" {
				continue
			}
			if dist[u.x][u.y]+1 <= dist[tx][ty] {
				dist[tx][ty] = dist[u.x][u.y] + 1
				q.push(point{tx, ty})
			}
		}
	}

	return dist
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var X int
	fmt.Fscan(reader, &X)
	var Y int
	fmt.Fscan(reader, &Y)
	var x = make([]int, n)
	var y = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		fmt.Fscan(reader, &y[i])
	}
	base := 250
	h := 500
	w := 500
	//base := 5
	//h := 10
	//w := 10

	s := make([][]string, h)
	for i := 0; i < h; i++ {
		s[i] = make([]string, w)
		for j := 0; j < w; j++ {
			s[i][j] = "."
		}
	}
	for i := 0; i < h; i++ {
		s[i][0] = "#"
		s[i][w-1] = "#"
	}
	for i := 0; i < w; i++ {
		s[0][i] = "#"
		s[h-1][i] = "#"
	}
	for i := 0; i < n; i++ {
		s[h-(y[i]+base)][x[i]+base] = "#"
	}
	//for i := 0; i < h; i++ {
	//	fmt.Fprintf(writer, "%v\n", s[i])
	//}

	dist := bfs2MinDist(h, w, h-base, base, s)
	//printer(dist, 1<<60)
	if dist[h-(Y+base)][X+base] == 1<<60 {
		fmt.Fprintf(writer, "%v\n", -1)
	} else {
		fmt.Fprintf(writer, "%v\n", dist[h-(Y+base)][X+base])
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
