package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var s string
	fmt.Fscan(reader, &s)

	s = strings.ReplaceAll(s, "p", "")
	s = strings.ReplaceAll(s, "s", "")
	s = strings.ReplaceAll(s, "t", "")

	for i := 0; i < len(s); i++ {
		if s[i] == "o"[0] {
			fmt.Fprintf(writer, "%v\n", i+1)
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "none")
}
