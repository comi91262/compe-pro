package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	var p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}

	m := map[int]int{}

	for i := 0; i < n; i++ {
		switch {
		case p[i] <= a:
			m[0]++
		case a < p[i] && p[i] <= b:
			m[1]++
		case b < p[i]:
			m[2]++
		}
	}

	fmt.Fprintf(writer, "%v\n", min(m[0], m[1], m[2]))
}
