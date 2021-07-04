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

	var a, b int
	fmt.Fscan(reader, &a, &b)

	if a <= b && b <= a*6 {
		fmt.Fprintf(writer, "Yes\n")
	} else {
		fmt.Fprintf(writer, "No\n")
	}
}
