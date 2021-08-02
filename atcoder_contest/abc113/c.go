package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var p = make([]int, m)
	var y = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &p[i])
		fmt.Fscan(reader, &y[i])
	}

	var g = make([][]int, n+1)
	for i := 0; i < m; i++ {
		g[p[i]] = append(g[p[i]], y[i])
	}

	for i := 1; i <= n; i++ {
		sort.Ints(g[i])
	}

	for i := 0; i < m; i++ {
		idx := binarySearch(g[p[i]], y[i], lowerBound)

		fmt.Fprintf(writer, "%06d%06d\n", p[i], idx+1)
	}
}
