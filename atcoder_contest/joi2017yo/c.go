package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func scanInt() int       { sc.Scan(); return parseInt(sc.Bytes()) }
func scanString() string { sc.Scan(); return string(sc.Bytes()) }
func scanFloat() float64 { sc.Scan(); return parseFloat(sc.Bytes()) }
func scanInts(n int) (ints []int) {
	ints = make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = scanInt()
	}
	return
}
func scanStrings(n int) (st []string) {
	st = make([]string, n)
	for i := 0; i < n; i++ {
		st[i] = scanString()
	}
	return
}
func scanFloats(n int) (fs []float64) {
	fs = make([]float64, n)
	for i := 0; i < n; i++ {
		fs[i] = scanFloat()
	}
	return fs
}

func parseInt(b []byte) (n int) {
	for _, ch := range b {
		ch -= '0'
		if ch <= 9 {
			n = n*10 + int(ch)
		}
	}
	if b[0] == '-' {
		return -n
	}
	return
}

var float64pow10 = []float64{
	1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6,
	1e7, 1e8, 1e9, 1e10, 1e11, 1e12,
	1e13, 1e14, 1e15, 1e16, 1e17, 1e18,
	1e19, 1e20, 1e21, 1e22,
}

func parseFloat(b []byte) float64 {
	f, dot := 0.0, 0
	for i, ch := range b {
		if ch == '.' {
			dot = i + 1
			continue
		}
		if ch -= '0'; ch <= 9 {
			f = f*10 + float64(ch)
		}
	}
	if dot != 0 {
		f /= float64pow10[len(b)-dot]
	}
	if b[0] == '-' {
		return -f
	}
	return f
}

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

var ans int
var d int

func dfs1(h, w, depth, x, y int, s [][]string, used [][]bool) {
	if d == depth {
		ans++
		return
	}
	if used[x][y] {
		return
	}

	var dx = [2]int{0}
	var dy = [2]int{1}

	used[x][y] = true
	for i := 0; i < 1; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || ny < 0 || nx >= h || ny >= w {
			continue
		}
		if s[nx][ny] == "#" {
			continue
		}
		dfs1(h, w, depth+1, nx, ny, s, used)
	}
	used[x][y] = false
}
func dfs2(h, w, depth, x, y int, s [][]string, used [][]bool) {
	if d == depth {
		ans++
		return
	}
	if used[x][y] {
		return
	}

	var dx = [1]int{1}
	var dy = [1]int{0}

	used[x][y] = true
	for i := 0; i < 1; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || ny < 0 || nx >= h || ny >= w {
			continue
		}
		if s[nx][ny] == "#" {
			continue
		}
		dfs2(h, w, depth+1, nx, ny, s, used)
	}
	used[x][y] = false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)
	h := scanInt()
	w := scanInt()
	d = scanInt()

	s := make([][]string, h)
	for i := 0; i < h; i++ {
		s[i] = strings.Split(scanString(), "")
	}

	used := make([][]bool, h)
	for i := 0; i < h; i++ {
		used[i] = make([]bool, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "#" {
				continue
			}
			dfs1(h, w, 1, i, j, s, used)
			dfs2(h, w, 1, i, j, s, used)
			// used[i][j] = true
		}
	}
	fmt.Fprintf(wr, "%v\n", ans)
}
