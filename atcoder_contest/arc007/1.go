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
	var x string
	fmt.Fscan(reader, &x)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	for i := range s {
		if s[i] != x {
			fmt.Fprintf(writer, "%v", s[i])
		}
	}
	fmt.Fprintf(writer, "\n")
}
