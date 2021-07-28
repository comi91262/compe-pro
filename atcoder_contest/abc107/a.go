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
	var i int
	fmt.Fscan(reader, &i)
	fmt.Fprintf(writer, "%v\n", n-i+1)
}
