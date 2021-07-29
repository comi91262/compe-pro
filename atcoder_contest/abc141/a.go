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
	case "Sunny":
		fmt.Fprintf(writer, "%v\n", "Cloudy")
	case "Cloudy":
		fmt.Fprintf(writer, "%v\n", "Rainy")
	case "Rainy":
		fmt.Fprintf(writer, "%v\n", "Sunny")
	}
}
