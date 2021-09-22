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

	var x int
	fmt.Fscan(reader, &x)
	switch {
	case 0 <= x && x < 40:
		fmt.Fprintf(writer, "%v\n", 40-x)
	case 40 <= x && x < 70:
		fmt.Fprintf(writer, "%v\n", 70-x)
	case 70 <= x && x < 90:
		fmt.Fprintf(writer, "%v\n", 90-x)
	case 90 <= x:
		fmt.Fprintf(writer, "%v\n", "expert")
	}
}
