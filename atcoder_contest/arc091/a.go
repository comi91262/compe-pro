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

	switch {
	case n == 1 && m == 1:
		fmt.Fprintf(writer, "%v\n", 1)
	case n == 1:
		fmt.Fprintf(writer, "%v\n", m-2)
	case m == 1:
		fmt.Fprintf(writer, "%v\n", n-2)
	default:
		fmt.Fprintf(writer, "%v\n", (n-2)*(m-2))
	}

}
