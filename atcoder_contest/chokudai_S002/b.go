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
	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a)
		fmt.Fscan(reader, &b)
		fmt.Fprintf(writer, "%v\n", a%b)

	}
}
