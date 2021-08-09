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

	if s[0] == s[1] && s[1] == s[2] && s[2] == s[3] {
		fmt.Fprintf(writer, "%v\n", "SAME")
	} else {
		fmt.Fprintf(writer, "%v\n", "DIFFERENT")
	}
}
