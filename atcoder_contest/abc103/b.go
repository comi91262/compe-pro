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
	var t string
	fmt.Fscan(reader, &t)

	for i := 0; i < len(s); i++ {
		if s == t {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}

		s = string(s[len(s)-1]) + s[:len(s)-1]
	}

	fmt.Fprintf(writer, "%v\n", "No")
}
