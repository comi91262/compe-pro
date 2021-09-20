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

var dx = [4]int{0, 1, 0, -1}
var dy = [4]int{1, 0, -1, 0}
var cnt int
var ans []point

func dfs2d(h, w, x, y int, s [][]string, acc []point, used [][]bool) {
	if len(acc) == cnt {
		ans = acc
		return
	}

	used[x][y] = true
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || ny < 0 || nx >= h || ny >= w {
			continue
		}
		if s[nx][ny] == "." {
			continue
		}
		if used[nx][ny] {
			continue
		}
		acc = append(acc, point{nx, ny})
		dfs2d(h, w, nx, ny, s, acc, used)
		acc = acc[:len(acc)-1]
	}
	used[x][y] = false
}

func main() {
	defer writer.Flush()
	var h, w int
	fmt.Fscan(reader, &h, &w)
	s := make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "#" {
				cnt++
			}
		}
	}

	used := make([][]bool, h)
	for i := 0; i < h; i++ {
		used[i] = make([]bool, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "." {
				continue
			}
			dfs2d(h, w, i, j, s, []point{{i, j}}, used)
			if len(ans) == cnt {
				goto L
			}
		}
	}

L:
	fmt.Fprintf(writer, "%v\n", cnt)
	for i := range ans {
		fmt.Fprintf(writer, "%v %v\n", ans[i].x+1, ans[i].y+1)
	}
}
