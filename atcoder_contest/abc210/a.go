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
	var n, a, x, y int
	fmt.Fscan(reader, &n, &a, &x, &y)

	sum := 0
	if a < n {
		sum = a*x + (n-a)*y
	} else {
		sum = n * x
	}

	fmt.Fprintf(writer, "%v\n", sum)
}
