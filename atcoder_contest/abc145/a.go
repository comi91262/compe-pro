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

	var r int
	fmt.Fscan(reader, &r)

    fmt.Fprintf(writer, "%v\n", r * r)
}
