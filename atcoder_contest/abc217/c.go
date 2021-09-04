package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}

	q := make([]int, n)
	for i := 0; i < n; i++ {
		q[p[i]-1] = i + 1
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v", q[i])
		if i != n-1 {
			fmt.Fprintf(writer, " ")
		}
	}
	fmt.Fprintf(writer, "\n")
}
