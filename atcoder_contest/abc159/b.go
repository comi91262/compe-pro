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

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}

	a := s[:len(s)/2]
	b := s[(len(s)+1)/2:]

	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[len(a)-1-i] {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}

	for i := 0; i < len(b)/2; i++ {
		if b[i] != b[len(b)-1-i] {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "Yes")
}
