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

func main() {
	defer writer.Flush()
	var c int
	fmt.Fscan(reader, &c)
	var n = make([]int, c)
	var m = make([]int, c)
	var l = make([]int, c)
	for i := 0; i < c; i++ {
		fmt.Fscan(reader, &n[i])
		fmt.Fscan(reader, &m[i])
		fmt.Fscan(reader, &l[i])
		d := []int{n[i], m[i], l[i]}
		sort.Slice(d, func(i, j int) bool { return d[i] > d[j] })
		n[i], m[i], l[i] = d[0], d[1], d[2]
	}

	s, t, u := 0, 0, 0
	for i := 0; i < c; i++ {
		s = max(s, n[i])
		t = max(t, m[i])
		u = max(u, l[i])
	}
	fmt.Fprintf(writer, "%v\n", s*t*u)
}
