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
	var t string
	fmt.Fscan(reader, &t)

	if len(s) > len(t) {
		fmt.Fprintf(writer, "%v\n", s)
	} else {
		fmt.Fprintf(writer, "%v\n", t)
	}
}
