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

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
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

	tot := sum(a)
	b := []int{}
	for i := 0; i < n; i++ {
		for a[i] > 0 {
			b = append(b, a[i]-a[i]/2)
			a[i] /= 2
		}
	}
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	for i := 0; i < min(m, len(b)); i++ {
		tot -= b[i]
	}
	fmt.Fprintf(writer, "%v\n", tot)
}
