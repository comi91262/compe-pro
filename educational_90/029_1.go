package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	defer writer.Flush()

	var w, n int
	fmt.Fscan(reader, &w, &n)

	var l = make([]int, n)
	var r = make([]int, n)
	var p = make([]int, n*2)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &l[i], &r[i])
		p[2*i] = l[i]
		p[2*i+1] = r[i]
	}

	sort.Ints(p)
	p = unique(p)

	for i := 0; i < n; i++ {
		l[i] = binarySearch(p, l[i], lowerBound)
		r[i] = binarySearch(p, r[i], lowerBound)
	}

	var q = make([]int, w)
	for i := 0; i < n; i++ {
		mx := 0
		for j := l[i]; j <= r[i]; j++ {
			mx = max(mx, q[j])
		}

		for j := l[i]; j <= r[i]; j++ {
			q[j] = mx + 1
		}

		fmt.Fprintf(writer, "%v\n", mx+1)
	}
}
