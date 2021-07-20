package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func contains(x int, a []int) bool {
	for _, e := range a {
		if x == e {
			return true
		}
	}
	return false
}

func main() {
	defer writer.Flush()
	var n, m int
	fmt.Fscan(reader, &n, &m)

	var a = make([]int, m)
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	var g = make([][]int, n+1)
	for i := 0; i < m; i++ {
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	var q int
	fmt.Fscan(reader, &q)

	var x = make([]int, q+1)
	var y = make([]int, q+1)
	for i := 1; i <= q; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}

	var color = make([]int, n+1)
	for i := 0; i <= n; i++ {
		color[i] = 1
	}
	var d = make([]int, n+1)
	var border = int(math.Sqrt(float64(2 * m)))

	// うにの集まりをつくる
	var e = []int{}
	for node, edges := range g {
		if len(edges) >= border {
			e = append(e, node)
		}
	}

	// うにの中から、gグラフ上で連結してるものをグラフhでもつ
	var h = make([][]int, n+1)
	for _, i := range e {
		for _, j := range e {
			if i == j || !contains(j, g[i]) {
				continue
			}
			h[i] = append(h[i], j)
		}
	}

	for i := 1; i <= q; i++ {
		ans := 1

		if contains(x[i], e) {
			ans = color[x[i]]
		} else {
			candidate := []int{d[x[i]]}
			for _, node := range g[x[i]] {
				candidate = append(candidate, d[node])
			}
			mx := max(candidate...)

			if mx != 0 {
				ans = y[mx]
			}
		}
		fmt.Fprintf(writer, "%v\n", ans)

		if contains(x[i], e) {
			color[x[i]] = y[i]
			for _, node := range h[x[i]] {
				color[node] = y[i]
			}
		} else {
			for _, node := range g[x[i]] {
				if contains(node, e) {
					color[node] = y[i]
				}
			}
		}

		d[x[i]] = i
	}
}
