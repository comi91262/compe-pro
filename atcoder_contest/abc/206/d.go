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

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var b = make([]int, n)
	var c = make([]int, n)
	for i := 0; i < n; i++ {
		if n%2 == 1 && i == n-1 {
			break
		}

		if i < n/2 {
			b[i] = a[i]
		} else {
			c[i-n/2] = a[n-1-i+n/2]
		}
	}

	fmt.Fprintf(writer, "%v %v\n", b, c)
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
