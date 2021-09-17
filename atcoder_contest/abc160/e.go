package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)

	var p = make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Fscan(reader, &p[i])
	}
	var q = make([]int, b)
	for i := 0; i < b; i++ {
		fmt.Fscan(reader, &q[i])
	}
	var r = make([]int, c)
	for i := 0; i < c; i++ {
		fmt.Fscan(reader, &r[i])
	}
	sort.Slice(p, func(i, j int) bool { return p[i] > p[j] })
	sort.Slice(q, func(i, j int) bool { return q[i] > q[j] })
	p = p[:x]
	q = q[:y]

	z := append(p, q...)
	z = append(z, r...)

	sort.Slice(z, func(i, j int) bool { return z[i] > z[j] })

	ans := 0
	for i := 0; i < x+y; i++ {
		ans += z[i]
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
