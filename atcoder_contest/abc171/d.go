package main

import (
	"bufio"
	"fmt"
	"os"
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

	var q int
	fmt.Fscan(reader, &q)
	var b = make([]int, q)
	var c = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &b[i])
		fmt.Fscan(reader, &c[i])
	}

	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]]++
	}
	total := sum(a)
	for i := 0; i < q; i++ {
		total -= m[b[i]] * b[i]
		total += m[b[i]] * c[i]
		fmt.Fprintf(writer, "%v\n", total)
		m[c[i]] += m[b[i]]
		m[b[i]] = 0
	}
}
