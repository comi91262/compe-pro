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
func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var h = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &h[i])
	}
	var w = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &w[i])
	}
	sort.Ints(h)

	if len(h) == 1 {
		mn := 1 << 60
		for i := 0; i < m; i++ {
			mn = min(abs(h[0]-w[i]), mn)
		}
		fmt.Fprintf(writer, "%v\n", mn)
		return
	}
	//fmt.Fprintf(writer, "%v\n", h)
	//fmt.Fprintf(writer, "%v\n", w)

	l := make([]int, n+1)
	r := make([]int, n+1)
	for i := 0; i < n-1; i += 2 {
		l[i+1] = l[i] + h[i+1] - h[i]
	}
	for i := 0; i < n-1; i += 2 {
		r[n-i-1] = r[n-i] + h[n-i-1] - h[n-i-2]
	}
	for i := 0; i < n; i++ {
		l[i+1] += l[i]
		r[n-i-1] += r[n-i]
	}

	//fmt.Fprintf(writer, "%v\n", l)
	//fmt.Fprintf(writer, "%v\n", r)

	ans := 1 << 60
	for i := 0; i < m; i++ {
		idx := sort.Search(n, func(j int) bool { return w[i] <= h[j] })
		//fmt.Fprintf(writer, "%v\n", idx)

		switch {
		case idx <= 1:
			ans = min(ans, abs(h[0]-w[i])+r[idx+1])
		case idx >= n-1:
			ans = min(ans, abs(h[n-1]-w[i])+l[idx-1])
		default:
			if idx%2 == 0 {
				ans = min(ans, abs(h[idx]-w[i])+l[idx-1]+r[idx+1])
			} else {
				ans = min(ans, abs(h[idx-1]-w[i])+l[idx-2]+r[idx])
			}
		}
		//fmt.Fprintf(writer, "%v\n", ans)
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
