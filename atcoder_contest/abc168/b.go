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
	var k int
	fmt.Fscan(reader, &k)
	var s string
	fmt.Fscan(reader, &s)

	if len(s) <= k {
		fmt.Fprintf(writer, "%v\n", s)
	} else {
		fmt.Fprintf(writer, "%v\n", s[:k]+"...")

	}

}
