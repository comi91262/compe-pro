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

	var s string
	fmt.Fscan(reader, &s)

	switch s {
	case "SUN":
		fmt.Fprintf(writer, "%v\n", 7)
	case "MON":
		fmt.Fprintf(writer, "%v\n", 6)
	case "TUE":
		fmt.Fprintf(writer, "%v\n", 5)
	case "WED":
		fmt.Fprintf(writer, "%v\n", 4)
	case "THU":
		fmt.Fprintf(writer, "%v\n", 3)
	case "FRI":
		fmt.Fprintf(writer, "%v\n", 2)
	case "SAT":
		fmt.Fprintf(writer, "%v\n", 1)
	}
}
