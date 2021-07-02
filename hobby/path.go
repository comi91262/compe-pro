package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var dx = [4]int{0, 1, 0, -1}
var dy = [4]int{1, 0, -1, 0}

type point struct {
	x int
	y int
}

var ans int

func dfs(x, y int, path []point, used [][]bool) {
	path = append(path, point{x: x, y: y})
	// fmt.Fprintf(writer, "%v %v\n", x, y)
	h := len(used)
	w := len(used[0])

	used[x][y] = true

	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || ny < 0 || nx >= h || ny >= w {
			continue
		}
		if used[nx][ny] {
			continue
		}
		dfs(nx, ny, path, used)
	}

	if len(path) == h*w {
		// fmt.Fprintf(writer, "%v\n", path)
		ans++
		// print(h, w, path)
	}

	path = path[0 : len(path)-1]
	used[x][y] = false

	// fmt.Fprintf(writer, "%v\n", used)
}

// +-+-+-+
// |○---+|
// +-+-+|+
// | | |||
// +-+-+|+
// | | |○|
// +-+-+-+

func print(h, w int, path []point) {
	m := make([][]int, h)
	for i := 0; i < h; i++ {
		m[i] = make([]int, w)
	}

	for i, p := range path {
		m[p.x][p.y] = i
	}

	bar := "+-+"
	for i := 2; i <= h; i++ {
		bar += "-+"
	}

	for i := 0; i < h; i++ {
		fmt.Fprintf(writer, "%v\n", bar)
		fmt.Fprintf(writer, "|")
		for j := 0; j < w; j++ {
			fmt.Fprintf(writer, "%v", m[i][j])
			fmt.Fprintf(writer, "|")
		}
		fmt.Fprintf(writer, "\n")
	}
	fmt.Fprintf(writer, "%v\n", bar)
}

func main() {
	defer writer.Flush()

	var h, w int
	fmt.Fscan(reader, &h, &w)

	c := make([][]int, h)
	for i := 0; i < h; i++ {
		c[i] = make([]int, w)
	}
	used := make([][]bool, h)
	for i := 0; i < h; i++ {
		used[i] = make([]bool, w)
	}
	path := make([]point, 0)
	dfs(0, 0, path, used)
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
// for j := 0; j < n; j++ {
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
// func min(arg ...int) int {
// 	min := arg[0]
// 	for _, x := range arg {
// 		if min > x {
// 			min = x
// 		}
// 	}
// 	return min
// }
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
