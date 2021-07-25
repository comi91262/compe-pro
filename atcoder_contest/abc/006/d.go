package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
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

func lis(a []int) []int {
	r := make([]int, len(a))

	b := make([]int, 0)
	for i := 0; i < len(a); i++ {
		cnt := binarySearch(b, a[i], lowerBound)

		if cnt == len(b) {
			b = append(b, a[i])
		} else {
			b[cnt] = a[i]
		}
		r[i] = cnt + 1
	}

	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	mx := 0
	d := lis(c)
	for i := 0; i < len(d); i++ {
		mx = max(mx, d[i])
	}
	fmt.Fprintf(writer, "%v\n", n-mx)
}
