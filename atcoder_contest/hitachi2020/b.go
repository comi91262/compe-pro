package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	first  int
	second int
}

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func minElement(arg ...int) int {
	min := arg[0]
	me := 0
	for i, x := range arg {
		if min > x {
			min = x
			me = i
		}
	}
	return me
}

func main() {
	defer writer.Flush()
	var A int
	fmt.Fscan(reader, &A)
	var B int
	fmt.Fscan(reader, &B)
	var M int
	fmt.Fscan(reader, &M)

	var a = make([]int, A)
	for i := 0; i < A; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var b = make([]int, B)
	for i := 0; i < B; i++ {
		fmt.Fscan(reader, &b[i])
	}

	ans := a[minElement(a...)] + b[minElement(b...)]
	for i := 0; i < M; i++ {
		var x int
		fmt.Fscan(reader, &x)
		var y int
		fmt.Fscan(reader, &y)
		x--
		y--
		var c int
		fmt.Fscan(reader, &c)

		ans = min(ans, a[x]+b[y]-c)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
