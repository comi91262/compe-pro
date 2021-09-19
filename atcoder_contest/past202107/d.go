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

	for i := 0; i < len(s)-2; i++ {
		if s[i] == "a" && s[i+1] == "x" && s[i+2] == "a" {
			s[i], s[i+1], s[i+2] = ".", ".", "."
			continue
		}
		if s[i] == "i" && s[i+1] == "x" && s[i+2] == "i" {
			s[i], s[i+1], s[i+2] = ".", ".", "."
			continue
		}
		if s[i] == "u" && s[i+1] == "x" && s[i+2] == "u" {
			s[i], s[i+1], s[i+2] = ".", ".", "."
			continue
		}
		if s[i] == "e" && s[i+1] == "x" && s[i+2] == "e" {
			s[i], s[i+1], s[i+2] = ".", ".", "."
			continue
		}
		if s[i] == "o" && s[i+1] == "x" && s[i+2] == "o" {
			s[i], s[i+1], s[i+2] = ".", ".", "."
			continue
		}
	}

	fmt.Fprintf(writer, "%v\n", strings.Join(s, ""))
}
