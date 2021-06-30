package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var dx = [4]int{0, 1, 0, -1}
var dy = [4]int{1, 0, -1, 0}
var c [][]string
var h, w int

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func dfs(sx, sy, x, y int, used [][]bool) int {
	if x == sx && y == sy && used[x][y] {
		return 0
	}

	used[x][y] = true

	path := -10000
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || ny < 0 || nx >= h || ny >= w || c[nx][ny] == "#" {
			continue
		}
		if (nx != sx || ny != sy) && used[nx][ny] {
			continue
		}

		r := dfs(sx, sy, nx, ny, used)
		path = max(path, r+1)
	}
	used[x][y] = false

	return path
}

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &h, &w)

	c = make([][]string, h)
	for i := 0; i < h; i++ {
		s := ""
		fmt.Fscan(reader, &s)
		c[i] = strings.Split(s, "")
	}

	used := make([][]bool, h)
	for i := 0; i < h; i++ {
		used[i] = make([]bool, w)
	}

	ans := -1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans = max(ans, dfs(i, j, i, j, used))
		}
	}

	if ans <= 2 {
		ans = -1
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
