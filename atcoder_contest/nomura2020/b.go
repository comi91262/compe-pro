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
	var ss string
	fmt.Fscan(reader, &ss)
	var t = strings.Split(ss, "")

	for i := range t {
		if t[i] == "?" {
			fmt.Fprintf(writer, "D")
		} else {
			fmt.Fprintf(writer, "%v", t[i])
		}
	}
	fmt.Fprintf(writer, "\n")
}

