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

	var n int
	fmt.Fscan(reader, &n)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	var k int
	fmt.Fscan(reader, &k)

	c := s[k-1]

	for i := 0; i < n; i++ {
		if s[i] != c {
			fmt.Fprintf(writer, "%v", "*")
		} else {
			fmt.Fprintf(writer, "%v", s[i])
		}
	}
	fmt.Fprintf(writer, "\n")
}
