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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var x = make([]int, n)
	var y = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}

	sort.Ints(x)
	sort.Ints(y)

	meanX := x[n/2]
	meanY := y[n/2]
	sum := 0
	for i := 0; i < n; i++ {
		sum += abs(x[i] - meanX)
		sum += abs(y[i] - meanY)
	}

	fmt.Fprintf(writer, "%d\n", sum)
}
