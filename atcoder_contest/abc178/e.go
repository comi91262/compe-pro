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
	var n int
	fmt.Fscan(reader, &n)
	var x = make([]int, n)
	var y = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}
	for i := 0; i < n; i++ {
		x[i], y[i] = x[i]-y[i], x[i]+y[i]
	}
	sort.Ints(x)
	sort.Ints(y)

	fmt.Fprintf(writer, "%v\n", max(x[n-1]-x[0], y[n-1]-y[0]))
}

