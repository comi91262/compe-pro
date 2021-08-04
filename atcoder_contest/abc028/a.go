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
	case n < 60:
		fmt.Fprintf(writer, "%v\n", "Bad")
	case 60 <= n && n < 90:
		fmt.Fprintf(writer, "%v\n", "Good")
	case 90 <= n && n < 100:
		fmt.Fprintf(writer, "%v\n", "Great")
	case n == 100:
		fmt.Fprintf(writer, "%v\n", "Perfect")
	}
}
