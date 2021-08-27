package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
func upperBound(value, boader int) bool {
	return boader < value
}

func binarySearch(a []int, boader int, criteria func(value, boader int) bool) int {
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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	var c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}
	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		idx := binarySearch(c, b[i], upperBound)
		if idx < n && b[i] < c[idx] {
			d[i] = n - idx
		}
	}
	for i := 0; i < n-1; i++ {
		d[i+1] += d[i]
	}

	ans := 0
	for i := 0; i < n; i++ {
		idx := binarySearch(b, a[i], upperBound)

		if idx < n && a[i] < b[idx] {
			ans += d[n-1]
			if idx != 0 {
				ans -= d[idx-1]
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
