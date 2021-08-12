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
	var X int
	fmt.Fscan(reader, &X)
	var Y int
	fmt.Fscan(reader, &Y)

	var x = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}
	var y = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &y[i])
	}

	sort.Ints(x)
	sort.Ints(y)

	for i := X + 1; i <= Y; i++ {
		if x[len(x)-1] < i && i <= y[0] {
			fmt.Fprintf(writer, "%v\n", "No War")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "War")

}
