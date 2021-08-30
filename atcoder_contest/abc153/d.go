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

	ans := 1
	for h > 0 {
		h /= 2
		ans *= 2
	}
	fmt.Fprintf(writer, "%v\n", ans-1)
}
