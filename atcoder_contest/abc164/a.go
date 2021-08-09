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

	var s int
	fmt.Fscan(reader, &s)
	var w int
	fmt.Fscan(reader, &w)

	if w >= s {
		fmt.Fprintf(writer, "%v\n", "unsafe")
	} else {
		fmt.Fprintf(writer, "%v\n", "safe")
	}
}
