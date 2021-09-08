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
	var w int
	fmt.Fscan(reader, &w)

	for i := 0; i < len(s); i++ {
		if i%w == 0 {
			fmt.Fprintf(writer, "%v", s[i])
		}
	}
	fmt.Fprintf(writer, "\n")
}
