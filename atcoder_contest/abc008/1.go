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
	var t int
	fmt.Fscan(reader, &t)
	fmt.Fprintf(writer, "%v\n", t-s+1)

}
