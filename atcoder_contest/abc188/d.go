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
func lowerBound(value, boader int) bool {
	return boader <= value
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
	var n int
	fmt.Fscan(reader, &n)
	var d int
	fmt.Fscan(reader, &d)

	var a = make([]int, n)
	var b = make([]int, n)
	var c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
		fmt.Fscan(reader, &c[i])
		b[i]++
	}
	e := []int{}
	for i := 0; i < n; i++ {
		e = append(e, a[i])
		e = append(e, b[i])
	}
	sort.Ints(e)
	e = unique(e)
	packed := map[int]int{}
	unpacked := map[int]int{}
	for i := 0; i < n; i++ {
		packed[a[i]] = binarySearch(e, a[i], lowerBound)
		packed[b[i]] = binarySearch(e, b[i], lowerBound)
		unpacked[packed[a[i]]] = a[i]
		unpacked[packed[b[i]]] = b[i]
	}

	f := make([]int, 2*n+1)
	for i := 0; i < n; i++ {
		f[packed[a[i]]] += c[i]
		f[packed[b[i]]] -= c[i]
	}
	for i := 0; i < 2*n; i++ {
		f[i+1] += f[i]
	}
	for i := 0; i < 2*n; i++ {
		if f[i] >= d {
			f[i] = d
		}
	}

	ans := 0
	for i := 0; i < 2*n; i++ {
		ans += f[i] * (unpacked[i+1] - unpacked[i])
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
