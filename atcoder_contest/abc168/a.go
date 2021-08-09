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

	switch n % 10 {
	case 0:
		fmt.Fprintf(writer, "%v\n", "pon")
	case 1:
		fmt.Fprintf(writer, "%v\n", "pon")
	case 2:
		fmt.Fprintf(writer, "%v\n", "hon")
	case 3:
		fmt.Fprintf(writer, "%v\n", "bon")
	case 4:
		fmt.Fprintf(writer, "%v\n", "hon")
	case 5:
		fmt.Fprintf(writer, "%v\n", "hon")
	case 6:
		fmt.Fprintf(writer, "%v\n", "pon")
	case 7:
		fmt.Fprintf(writer, "%v\n", "hon")
	case 8:
		fmt.Fprintf(writer, "%v\n", "pon")
	case 9:
		fmt.Fprintf(writer, "%v\n", "hon")
	}
}
