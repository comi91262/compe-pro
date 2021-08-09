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

	ans := 0

	if s[0] == "R" || s[1] == "R" || s[2] == "R" {
		ans++
	}

	if s[0] == s[1] && s[1] == "R" {
		ans++
	}
	if s[1] == s[2] && s[1] == "R" {
		ans++
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
