package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func isOK(a *[]int, index, key int) bool {
	return (*a)[index] >= key
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

	var c = make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			c[i] = a[i] - 1
			continue
		}
		c[i] = c[i-1] + a[i] - a[i-1] - 1
	}

	var k = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &k[i])
	}

	for i := 0; i < q; i++ {
		ans := 0
		if c[n-1] < k[i] {
			ans = k[i] - c[n-1] + a[n-1]
		} else {
			idx := binarySearch(&c, k[i])
			ans = a[idx] - 1 - (c[idx] - k[i])
		}
		fmt.Fprintf(writer, "%d\n", ans)
	}
}
