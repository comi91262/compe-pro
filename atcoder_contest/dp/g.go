package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func scanInt() int       { sc.Scan(); return parseInt(sc.Bytes()) }
func scanString() string { sc.Scan(); return string(sc.Bytes()) }
func scanInts(n int) (ints []int) {
	ints = make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = scanInt()
	}
	return
}

func scanPairInts(n int) (a, b []int) {
	a = make([]int, n)
	b = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
		b[i] = scanInt()
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

func chmax(x *int, y int) {
	*x = max(*x, y)
}

type edge struct {
	to   int
	cost int
}

func dfs(x int, dp []int, used []bool, g [][]edge) (path int) {
	if used[x] {
		return dp[x]
	}

	used[x] = true
	for _, y := range g[x] {
		chmax(&path, dfs(y.to, dp, used, g)+1)
	}
	dp[x] = path
	return
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)

	n := scanInt()
	m := scanInt()

	g := make([][]edge, n)
	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = scanInt()
		b[i] = scanInt()
		a[i]--
		b[i]--
		g[a[i]] = append(g[a[i]], edge{b[i], 0})
	}

	dp := make([]int, n)
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		dfs(i, dp, used, g)
	}
	ans := 0
	for i := range dp {
		chmax(&ans, dp[i])
	}
	fmt.Fprintf(wr, "%v\n", ans)

}
