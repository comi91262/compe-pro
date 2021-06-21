package main

import (
	"bufio"
	"fmt"
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

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
}

func main() {
	defer writer.Flush()

	var h, w int
	fmt.Fscan(reader, &h, &w)

	var p = make([][]int, h)
	for i := 0; i < h; i++ {
		temp := make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Fscan(reader, &temp[j])
		}
		p[i] = temp
	}

	var m = pow(2, h)

	ans := 0
	for i := 0; i < m; i++ {
		if i == 0 {
			continue
		}

		q1 := make([][]int, 0, m)
		for j := 0; j < h; j++ {
			if i&(1<<j) != 0 {
				q1 = append(q1, p[j])
			}
		}

		q2 := make([]int, w)
		for j := 0; j < w; j++ {
			pre := 0
			check := true
			for k := 0; k < len(q1); k++ {
				if q1[k][j] != pre && pre != 0 {
					check = false
					break
				}
				pre = q1[k][j]
			}

			if check {
				q2[j] = q1[0][j]
			}
		}

		q3 := map[int]int{}
		for j := 0; j < w; j++ {
			q3[q2[j]]++
		}

		q4 := 0
		for i, v := range q3 {
			if i == 0 {
				continue
			}
			q4 = max(q4, v*len(q1))
		}
		ans = max(ans, q4)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}

// var n, q int
// fmt.Fscan(reader, &n, &q)

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
// func max(x, y int) int {
// 	return int(math.Max(float64(x), float64(y)))
// }

// var g = make([][]int, n+1)
// var a = make([]int, n+1)
// var b = make([]int, n+1)
// for i := 1; i < n; i++ {
// 	fmt.Fscan(reader, &a[i], &b[i])
// 	g[a[i]] = append(g[a[i]], b[i])
// 	g[b[i]] = append(g[b[i]], a[i])
// }

// m := map[int]int{}
// a := make([]int, n)
// for i := 0; i < len(a); i++ {
// 	m[a[i]]++
// }
