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
	fmt.Fprintf(writer, "%v\n", n)
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", 1)
	}
}
