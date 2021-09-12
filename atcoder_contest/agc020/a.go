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
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	if (a-b+2)%2 == 0 {
		fmt.Fprintf(writer, "%v\n", "Alice")
	} else {
		fmt.Fprintf(writer, "%v\n", "Borys")
	}
}
