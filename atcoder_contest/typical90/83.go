package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func lowerBound(value, boader int) bool {
	return boader <= value
}

func binarySearch(a []int, boader int, criteria func(value, boader int) bool) int {
	abs := func(x int) int {
		if x >= 0 {
			return x
		} else {
			return x * -1
		}
	}

	ng := -1
	ok := len(a)

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if criteria(a[mid], boader) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
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

func contains(x int, a []int) bool {
	// idx := binarySearch(a, x, lowerBound)
	// return (idx < len(a)) && a[idx] == x
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

	// constrains: increment order
	var e = []int{}
	for node, edges := range g {
		if len(edges) >= border {
			e = append(e, node)
		}
	}

	var h = make([][]int, n+1)
	for _, i := range e {
		for _, j := range e {
			if i == j {
				continue
			}
			flag := false
			for _, k := range g[i] {
				if k == j {
					flag = true
					break
				}
			}
			if !flag {
				continue
			}

			h[i] = append(h[i], j)
			// fmt.Fprintf(writer, "%v %v\n", i, j)
		}
	}
	cnt := 0
	for i := 1; i < n+1; i++ {
		if len(h[i]) > 0 {
			cnt++
		}
	}
	// fmt.Fprintf(writer, "%v\n", cnt)
	// fmt.Fprintf(writer, "%v\n", e)

	for i := 1; i <= q; i++ {
		ans := 1
		// fmt.Fprintf(writer, "%v %v\n", x[i], y[i])

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
				// if contains(x[mx], e) {
				// 	ans = color[x[mx]]
				// } else {
				// 	ans = y[mx]
				// 	//	for _, node := range g[x[mx]] {
				// 	//		if contains(node, e) {
				// 	//			color[node] = y[mx]
				// 	//		}
				// 	//	}
				// }
			}
		}
		fmt.Fprintf(writer, "%v\n", ans)

		if contains(x[i], e) {
			color[x[i]] = y[i]
			for _, node := range h[x[i]] {
				// fmt.Fprintf(writer, "%v\n", node)
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
		// fmt.Fprintf(writer, "%v\n", color)
		// fmt.Fprintf(writer, "%v\n", ans)
	}

	// for i := 1; i <= n; i++ {
	// 	if !contains(i, e) {
	// 		color[i] = x[d[i]]
	// 	}
	// }

}
