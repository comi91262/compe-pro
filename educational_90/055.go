package main

import (
	"bufio"
	"fmt"
	"os"
)

var p, q int

func prod(a []int, n, total int) int {
	if n == 0 {
		if total%p == q {
			return 1
		}
		return 0
	}

	if len(a) < n {
		return 0
	}

	return prod(a[1:], n-1, total*a[0]) + prod(a[1:], n, total)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n, &p, &q)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	for i := 0; i < n; i++ {
		a[i] = a[i] % p
	}

	fmt.Fprintf(writer, "%d\n", prod(a, 5, 1))
}

// var n int
// fmt.Fscan(reader, &n)
// fmt.Fprintf(writer, "%d\n", n)

//	var n int
//	fmt.Fscan(reader, &n)
//	var a = make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Fscan(reader, &a[i])
//	}
// snip
// for i := 0; i < n; i++ {
// for j := 0; j < n; j++ {
// }
// }
// func abs(x int) int {
// 	return int(math.Abs(float64(x)))
// }
// func min(x, y int) int {
// 	return int(math.Min(float64(x), float64(y)))
// }
