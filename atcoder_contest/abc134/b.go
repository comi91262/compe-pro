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
	var d int
	fmt.Fscan(reader, &d)

	fmt.Fprintf(writer, "%v\n", (n+2*d)/(2*d+1))

}
