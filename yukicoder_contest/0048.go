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

	var a, b int
	fmt.Fscan(reader, &a)
	fmt.Fscan(reader, &b)

	if b%a == 0 {
		fmt.Fprintf(writer, "%v\n", b/a)
	} else {
		fmt.Fprintf(writer, "%v\n", b/a+1)
	}
}
