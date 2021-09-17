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
type queue []point

func (q *queue) push(n point) {
	*q = append(*q, n)
}

func (q *queue) pop() point {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) front() point {
	return (*q)[0]
}

func (q *queue) empty() bool {
	return len(*q) == 0
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

	q := queue{}
	dist[sx][sy] = 0
	q.push(point{sx, sy})

	for !q.empty() {
		u := q.pop()
		for i := 0; i < 2; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]

			if tx < 0 || tx >= h || ty < 0 || ty >= w {
				continue
			}
			if dist[tx][ty] != inf || s[tx][ty] == "." {
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

	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)
	s := make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	dist := bfs2MinDist(h, w, 0, 0, s)
	// printer(dist, 1<<60)

	m := map[int]int{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			m[dist[i][j]]++
			if s[i][j] == "#" && dist[i][j] == 1<<60 {
				fmt.Fprintf(writer, "%v\n", "Impossible")
				return
			}
		}
	}
	for k, v := range m {
		if k == 1<<60 {
			continue
		}
		if v > 1 {
			fmt.Fprintf(writer, "%v\n", "Impossible")
			return
		}

	}

	fmt.Fprintf(writer, "%v\n", "Possible")

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
