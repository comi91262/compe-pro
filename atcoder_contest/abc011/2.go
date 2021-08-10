package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)

	for i, r := range s {
		if i == 0 {
			fmt.Fprintf(writer, "%c", unicode.ToUpper(r))
			continue
		}
		fmt.Fprintf(writer, "%c", unicode.ToLower(r))
	}

	fmt.Fprintf(writer, "\n")
}
