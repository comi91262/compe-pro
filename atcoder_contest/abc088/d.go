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

func (q *queue) empty() bool {
	return len(*q) == 0
}

func bfs(h, w, sx, sy int, s [][]string) [][]int {
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
		for i := 0; i < 4; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]
			if 0 <= tx && tx < h && 0 <= ty && ty < w && dist[tx][ty] == inf && dist[u.x][u.y]+1 <= dist[tx][ty] && s[tx][ty] == "." {
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
		var ss string
		fmt.Fscan(reader, &ss)
		s[i] = strings.Split(ss, "")
	}

	b := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "#" {
				b++
			}
		}
	}

	path := bfs(h, w, 0, 0, s)[h-1][w-1]
	if path == 1<<60 {
		fmt.Fprintf(writer, "%v\n", -1)
	} else {
		fmt.Fprintf(writer, "%v\n", h*w-b-(path+1))
	}
}
