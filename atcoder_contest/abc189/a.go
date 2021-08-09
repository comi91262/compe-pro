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

	if s[0] == s[1] && s[1] == s[2] {
		fmt.Fprintf(writer, "%v\n", "Won")
	} else {
		fmt.Fprintf(writer, "%v\n", "Lost")
	}
}
