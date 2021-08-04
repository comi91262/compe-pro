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

	var h int
	fmt.Fscan(reader, &h)
	var a int
	fmt.Fscan(reader, &a)

	cnt := 0
	for h > 0 {
		cnt++
		h -= a
	}

	fmt.Fprintf(writer, "%v\n", cnt)
}
