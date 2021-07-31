package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

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

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}

	sort.Ints(a)
	sort.Ints(b)

	ans := 1 << 60
	for i := 0; i < m; i++ {
		idx := binarySearch(a, b[i], lowerBound)

		if idx < n {
			ans = min(ans, abs(a[idx]-b[i]))
		}
		if idx+1 < n {
			ans = min(ans, abs(a[idx+1]-b[i]))
		}
		if 0 < idx-1 {
			ans = min(ans, abs(a[idx-1]-b[i]))
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
