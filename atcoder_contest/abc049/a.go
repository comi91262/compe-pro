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

	var c string
	fmt.Fscan(reader, &c)

	switch c {
	case "a":
		fmt.Fprintf(writer, "%v\n", "vowel")
	case "i":
		fmt.Fprintf(writer, "%v\n", "vowel")
	case "u":
		fmt.Fprintf(writer, "%v\n", "vowel")
	case "e":
		fmt.Fprintf(writer, "%v\n", "vowel")
	case "o":
		fmt.Fprintf(writer, "%v\n", "vowel")
	default:
		fmt.Fprintf(writer, "%v\n", "consonant")
	}
}
