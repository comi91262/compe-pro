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
	x   int
	y   int
	pre int
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

func bfs(h, w int, sx, sy int, used [][]bool, s [][]string) int {
	const inf = 1 << 60
	var dx = [4]int{1, 0, -1, 0}
	var dy = [4]int{0, 1, 0, -1}

	q := queue{}
	used[sx][sy] = true
	q.push(point{sx, sy, 0})

	bl := 1
	wh := 0
	for !q.empty() {
		u := q.pop()
		for i := 0; i < 4; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]

			if 0 <= tx && tx < h && 0 <= ty && ty < w && !used[tx][ty] {
				if u.pre == 0 && s[tx][ty] == "." {
					wh++
					q.push(point{tx, ty, 1})
					used[tx][ty] = true
				} else if u.pre == 1 && s[tx][ty] == "#" {
					bl++
					q.push(point{tx, ty, 0})
					used[tx][ty] = true
				}
			}
		}
	}

	return wh * bl
}

func main() {
	defer writer.Flush()

	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	a := make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		a[i] = strings.Split(tmp, "")
	}

	used := make([][]bool, h)
	for i := 0; i < h; i++ {
		used[i] = make([]bool, w)
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if a[i][j] == "#" {
				ans += bfs(h, w, i, j, used, a)
			}
		}

	}
	fmt.Fprintf(writer, "%v\n", ans)
}
