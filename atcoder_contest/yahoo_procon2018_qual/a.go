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

	if s[0] == "y" && s[1] == "a" && s[2] == "h" && s[3] == s[4] {
		fmt.Fprintf(writer, "%v\n", "YES")
	} else {
		fmt.Fprintf(writer, "%v\n", "NO")
	}
}
