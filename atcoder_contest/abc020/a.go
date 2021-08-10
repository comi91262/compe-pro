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
	var q int
	fmt.Fscan(reader, &q)

	switch q {
	case 1:
		fmt.Fprintf(writer, "%v\n", "ABC")
	case 2:
		fmt.Fprintf(writer, "%v\n", "chokudai")
	}

}
