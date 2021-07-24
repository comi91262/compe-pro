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

	var a, b float64
	fmt.Fscan(reader, &a, &b)

	fmt.Fprintf(writer, "%v\n", (a-b)/3+b)
}
