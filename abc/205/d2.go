package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const inf = math.MaxInt64

func abs(x int) int {
	if x < 0 {
		return x * -1
	} else {
		return x
	}
}

func isOK(a *[]int, index, key int) bool {
	return (*a)[index] > key
}

func binarySearch(a *[]int, key int) int {
	ng := -1
	ok := len(*a)

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if isOK(a, mid, key) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func main() {
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	for i := 0; i < q; i++ {
		var k int
		fmt.Fscan(reader, &k)
		ng := 0
		ok := inf

		for abs(ok-ng) > 1 {
			mid := (ok + ng) / 2
			idx := binarySearch(&a, mid)

			if mid-idx >= k {
				ok = mid
			} else {
				ng = mid
			}
		}

		fmt.Fprintf(writer, "%d\n", ok)
	}

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
// const d = 1_000_000_000 + 7
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
// func max(x, y int) int {
// 	return int(math.Max(float64(x), float64(y)))
// }
