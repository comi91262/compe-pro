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
	var r int
	fmt.Fscan(reader, &r)

	if n >= 10 {
		fmt.Fprintf(writer, "%v\n", r)
	} else {
		fmt.Fprintf(writer, "%v\n", r + 100*(10-n))
	}
}
