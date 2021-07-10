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

	var n, q int
	fmt.Fscan(reader, &n, &q)

	if q < n {
		fmt.Fprintf(writer, "%v\n", 0)
	} else {
		fmt.Fprintf(writer, "%v\n", q-n+1)
	}
}
