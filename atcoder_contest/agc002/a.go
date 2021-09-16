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
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	if a <= 0 && 0 <= b {
		fmt.Fprintf(writer, "%v\n", "Zero")
		return
	}
	if a > 0 && b > 0 {
		fmt.Fprintf(writer, "%v\n", "Positive")
		return
	}

	if (a-b)%2 == 0 {
		fmt.Fprintf(writer, "%v\n", "Negative")
	} else {
		fmt.Fprintf(writer, "%v\n", "Positive")
	}
}
