package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func enumarate(size, init, inc int) []int {
	a := make([]int, size)
	if size == 0 {
		return a
	}

	a[0] = init
	for i := 1; i < size; i++ {
		a[i] = a[i-1] + inc
	}

	return a
}

func unique(a []int) []int {
	r := make([]int, 0)
	m := make(map[int]bool, 0)

	for _, e := range a {
		if !m[e] {
			m[e] = true
			r = append(r, e)
		}
	}

	return r
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

func compression(a []int) []int {
	var n = len(a)
	var p = make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = a[i]
	}

	sort.Ints(p)
	p = unique(p)

	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = binarySearch(p, a[i], lowerBound)
	}

	return r
}

func main() {
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	order := compression(a)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = k / n
	}

	for i := 0; i < n; i++ {
		if order[i] >= k%n {
			continue
		}
		ans[i]++
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", ans[i])
	}
}
