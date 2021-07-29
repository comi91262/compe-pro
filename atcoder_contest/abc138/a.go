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

	var a int
	fmt.Fscan(reader, &a)
	var s string
	fmt.Fscan(reader, &s)

	if a >= 3200 {
		fmt.Fprintf(writer, "%v\n", s)
	} else {
		fmt.Fprintf(writer, "%v\n", "red")
	}
}
