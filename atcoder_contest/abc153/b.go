package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}

func main() {
	defer writer.Flush()

	var h int
	fmt.Fscan(reader, &h)
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	if h-sum(a) > 0 {
		fmt.Fprintf(writer, "%v\n", "No")
	} else {
		fmt.Fprintf(writer, "%v\n", "Yes")
	}
}
