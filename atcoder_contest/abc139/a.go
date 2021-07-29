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

	cnt := 0
	if s[0] == t[0] {
		cnt++
	}
	if s[1] == t[1] {
		cnt++
	}
	if s[2] == t[2] {
		cnt++
	}

    fmt.Fprintf(writer, "%v\n", cnt)
}
