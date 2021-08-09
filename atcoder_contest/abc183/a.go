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

	var x int
	fmt.Fscan(reader, &x)

	if x >= 0 {
		fmt.Fprintf(writer, "%v\n", x)
	} else {
		fmt.Fprintf(writer, "%v\n", 0)
	}
}
