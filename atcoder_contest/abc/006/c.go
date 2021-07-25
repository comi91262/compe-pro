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

	var n, m int
	fmt.Fscan(reader, &n, &m)

	for c := 0; c <= n; c++ {
		a := c + 3*n - m
		b := m - 2*n - 2*c

		if a >= 0 && b >= 0 {
			fmt.Fprintf(writer, "%v %v %v\n", a, b, c)
			return
		}
	}

	fmt.Fprintf(writer, "-1 -1 -1\n")
}
