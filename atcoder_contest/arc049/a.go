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
	var s = strings.Split(ss, "")
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)
	var d int
	fmt.Fscan(reader, &d)

	for i := 0; i < len(s); i++ {
		if i == a || i == b || i == c || i == d {
			fmt.Fprintf(writer, "\"")
		}
		fmt.Fprintf(writer, "%v", s[i])
	}
	if len(s) == d {
		fmt.Fprintf(writer, "\"")
	}

	fmt.Fprintf(writer, "\n")
}
