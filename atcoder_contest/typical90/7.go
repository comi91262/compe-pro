package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
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
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var q int
	fmt.Fscan(reader, &q)
	var b = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &b[i])
	}

	sort.Ints(a)

	for i := 0; i < q; i++ {
		var m = binarySearch(&a, b[i])
		var v int
		if m == len(a) {
			v = abs(b[i] - a[len(a)-1])
		} else if m > 0 {
			v = min(abs(b[i]-a[m]), abs(b[i]-a[m-1]))
		} else {
			v = abs(b[i] - a[m])
		}
		fmt.Fprintf(writer, "%d\n", v)
	}
}
