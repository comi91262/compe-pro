package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const inf = 1 << 60

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func calcRange(li []int, x int) (l, r int) {
	n := len(li)

	r = sort.Search(n, func(k int) bool { return x <= li[k] })
	l = r - 1

	if r >= n {
		r = inf
	} else {
		r = li[r]
	}

	if l < 0 {
		l = -1
	} else {
		l = li[l]
	}

	return
}

func isOk(l, b [][]int, i, j int) bool {
	if len(b[i]) == 0 && len(l[i]) > 0 {
		return true
	}
	if len(l[i]) == 0 {
		return false
	}

	ll, lr := calcRange(l[i], j)
	bl, br := calcRange(b[i], j)

	if ll == -1 && j <= br && br <= lr {
		return false
	}
	if (ll < bl && bl <= j) && lr == inf {
		return false
	}
	if ll < bl && bl <= j && j <= br && br <= lr {
		return false
	}

	return true
}

func main() {
	defer writer.Flush()
	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	s := make([][]string, h)
	for i := 0; i < h; i++ {
		s[i] = make([]string, w)
		for j := 0; j < w; j++ {
			s[i][j] = "."
		}
	}
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(reader, &a)
		var b int
		fmt.Fscan(reader, &b)
		a--
		b--
		s[a][b] = "L"
	}
	for i := 0; i < m; i++ {
		var c int
		fmt.Fscan(reader, &c)
		var d int
		fmt.Fscan(reader, &d)
		c--
		d--
		s[c][d] = "#"
	}

	var rl = make([][]int, h)
	var rb = make([][]int, h)
	var cl = make([][]int, w)
	var cb = make([][]int, w)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "L" {
				rl[i] = append(rl[i], j)
			}
			if s[i][j] == "#" {
				rb[i] = append(rb[i], j)
			}
		}
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if s[j][i] == "L" {
				cl[i] = append(cl[i], j)
			}
			if s[j][i] == "#" {
				cb[i] = append(cb[i], j)
			}
		}
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "L" {
				ans++
				continue
			}
			if s[i][j] == "#" {
				continue
			}

			if !isOk(rl, rb, i, j) && !isOk(cl, cb, j, i) {
				continue
			}
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
