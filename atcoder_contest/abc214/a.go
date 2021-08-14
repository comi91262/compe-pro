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

	switch {
	case 1 <= n && n <= 125:
		fmt.Fprintf(writer, "%v\n", 4)
	case 126 <= n && n <= 211:
		fmt.Fprintf(writer, "%v\n", 6)
	case 212 <= n && n <= 214:
		fmt.Fprintf(writer, "%v\n", 8)
	}
}
