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
	var t int
	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		switch {
		case n%4 == 0:
			fmt.Fprintf(writer, "%v\n", "Even")
		case n%2 == 0:
			fmt.Fprintf(writer, "%v\n", "Same")
		default:
			fmt.Fprintf(writer, "%v\n", "Odd")
		}
	}

}
