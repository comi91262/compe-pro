package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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
	ok := -1
	ng := len(a)

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

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)

	d := []int{a, b, c}
	sort.Slice(d, func(i, j int) bool { return d[i] > d[j] })

	fmt.Fprintf(writer, "%v\n", binarySearch(d, a, lowerBound)+1)
	fmt.Fprintf(writer, "%v\n", binarySearch(d, b, lowerBound)+1)
	fmt.Fprintf(writer, "%v\n", binarySearch(d, c, lowerBound)+1)

}
