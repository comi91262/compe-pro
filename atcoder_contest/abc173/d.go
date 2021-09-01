package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })

	b := make([]int, n-1)
	for i := 0; i < (n+1)/2; i++ {
		if i == 0 {
			b[i] = a[i]
			continue
		}
		b[2*i-1] = a[i]
		if 2*i < n-1 {
			b[2*i] = a[i]
		}
	}
	fmt.Fprintf(writer, "%v\n", sum(b))
}
