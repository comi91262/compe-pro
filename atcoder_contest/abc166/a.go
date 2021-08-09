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

	var s string
	fmt.Fscan(reader, &s)

	if s == "ARC" {
		fmt.Fprintf(writer, "%v\n", "ABC")
	} else {
		fmt.Fprintf(writer, "%v\n", "ARC")
	}
}
