package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type state struct {
	x int
	y int
}

type queue []state

func (q *queue) push(n state) {
	*q = append(*q, n)
}

func (q *queue) pop() state {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func bfs2(h, w int, sx, sy, ex, ey int, dist [][]int, s [][]string) {
	const inf = 1 << 60
	var dx = [4]int{1, 0, -1, 0}
	var dy = [4]int{0, 1, 0, -1}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			dist[i][j] = inf
		}
	}

	q := queue{}
	dist[sx][sy] = 0
	q.push(state{sx, sy})

	for !q.empty() {
		u := q.pop()
		for i := 0; i < 4; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]
			if 0 <= tx && tx < h && 0 <= ty && ty < w && s[tx][ty] == "." && dist[tx][ty] == inf {
				dist[tx][ty] = dist[u.x][u.y] + 1
				q.push(state{tx, ty})
			}
		}
	}
}

func main() {
	defer writer.Flush()

	var h, w int
	fmt.Fscan(reader, &h, &w)

	var sx, sy, tx, ty int
	fmt.Fscan(reader, &sx, &sy, &tx, &ty)
	sx--
	sy--
	tx--
	ty--

	var s = make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	dist := make([][]int, h)
	for i := 0; i < h; i++ {
		dist[i] = make([]int, w)
	}
	bfs2(h, w, sx, sy, tx, ty, dist, s)

	fmt.Fprintf(writer, "%v\n", dist[tx][ty])
}
