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
	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	cnt := 0
	for i := 0; i < n; i++ {
		if h <= a[i] && w <= b[i] {
			cnt++
		}
	}
	fmt.Fprintf(writer, "%v\n", cnt)
}
