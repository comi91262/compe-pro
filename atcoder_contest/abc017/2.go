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

	for {
		switch {
		case strings.HasPrefix(s, "ch"):
			s = s[2:]
		case strings.HasPrefix(s, "o"):
			s = s[1:]
		case strings.HasPrefix(s, "k"):
			s = s[1:]
		case strings.HasPrefix(s, "u"):
			s = s[1:]
		case s == "":
			fmt.Fprintf(writer, "%v\n", "YES")
			return
		default:
			fmt.Fprintf(writer, "%v\n", "NO")
			return
		}
	}

}
