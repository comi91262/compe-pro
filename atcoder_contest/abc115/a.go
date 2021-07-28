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
	var d int
	fmt.Fscan(reader, &d)

	switch d {
	case 25:
		fmt.Fprintf(writer, "%v\n", "Christmas")
	case 24:
		fmt.Fprintf(writer, "%v\n", "Christmas Eve")
	case 23:
		fmt.Fprintf(writer, "%v\n", "Christmas Eve Eve")
	case 22:
		fmt.Fprintf(writer, "%v\n", "Christmas Eve Eve Eve")
	}
}
