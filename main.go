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
