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

	switch {
	case r < 1200:
		fmt.Fprintf(writer, "%v\n", "ABC")
	case r < 2800:
		fmt.Fprintf(writer, "%v\n", "ARC")
	default:
		fmt.Fprintf(writer, "%v\n", "AGC")
	}

}
