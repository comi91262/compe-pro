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

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	var x = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &x[i])
	}

	if n >= m {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	sort.Ints(x)
	var y = make([]int, m-1)
	for i := 0; i < m-1; i++ {
		y[i] = x[i+1] - x[i]
	}
	sort.Slice(y, func(i, j int) bool { return y[i] > y[j] })

	// fmt.Fprintf(writer, "%v\n", y)
	ans := 0
	for i := n - 1; i < m-1; i++ {
		ans += y[i]
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
