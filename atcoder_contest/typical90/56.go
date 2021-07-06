package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, s int
	fmt.Fscan(reader, &n, &s)

	var a = make([]int, n)
	var b = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	var dp = make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, s+1)
	}

	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < s+1; j++ {
			if j-a[i] >= 0 && dp[i][j-a[i]] {
				dp[i+1][j] = true
			}
			if j-b[i] >= 0 && dp[i][j-b[i]] {
				dp[i+1][j] = true
			}
		}
	}

	if !dp[n][s] {
		fmt.Fprintf(writer, "Impossible")
		return
	}

	r := ""
	for n > 0 {
		if s-a[n-1] >= 0 && dp[n-1][s-a[n-1]] {
			r = "A" + r

			n = n - 1
			s = s - a[n]
		} else if s-b[n-1] >= 0 && dp[n-1][s-b[n-1]] {
			r = "B" + r
			n = n - 1
			s = s - b[n]
		}
	}

	fmt.Fprintf(writer, "%v", r)
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
// fmt.Fprintf(writer, "%v\n", n)

//	var n int
//	fmt.Fscan(reader, &n)
//	var a = make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Fscan(reader, &a[i])
//	}
// const d = 1_000_000_000 + 7
// for i := 0; i < n; i++ {
// for j := 0; j < n; j++ {
// for k := 0; k < n; k++ {
// }
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
