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

	var l int
	fmt.Fscan(reader, &l)
	var h int
	fmt.Fscan(reader, &h)
	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	for i := 0; i < n; i++ {
		switch {
		case a[i] < l:
			fmt.Fprintf(writer, "%v\n", l-a[i])
		case l <= a[i] && a[i] <= h:
			fmt.Fprintf(writer, "%v\n", 0)
		case h < a[i]:
			fmt.Fprintf(writer, "%v\n", -1)
		}
	}
}
