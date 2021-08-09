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
	if s[len(s)-1] == "s"[0] {
		fmt.Fprintf(writer, "%v\n", s+"es")
	} else {
		fmt.Fprintf(writer, "%v\n", s+"s")
	}
}
