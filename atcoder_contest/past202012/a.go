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

	for i := 0; i < len(s)-2; i++ {
		if s[i] == "o" && s[i] == s[i+1] && s[i] == s[i+2] {
			fmt.Fprintf(writer, "%v\n", "o")
			return
		}
		if s[i] == "x" && s[i] == s[i+1] && s[i] == s[i+2] {
			fmt.Fprintf(writer, "%v\n", "x")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "draw")
}
