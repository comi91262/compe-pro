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
	var m int
	fmt.Fscan(reader, &m)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &c[i])
	}

	for i := 0; i < m; i++ {
		if n <= a {
			n += b
		}
		if n-c[i] < 0 {
			fmt.Fprintf(writer, "%v\n", i+1)
			return
		}
		n -= c[i]
	}
	fmt.Fprintf(writer, "%v\n", "complete")
}
