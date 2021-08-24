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
	var s int
	fmt.Fscan(reader, &s)
	var d int
	fmt.Fscan(reader, &d)

	var x = make([]int, n)
	var y = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		fmt.Fscan(reader, &y[i])
	}

	for i := 0; i < n; i++ {
		if x[i] < s && y[i] > d {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "No")
}
