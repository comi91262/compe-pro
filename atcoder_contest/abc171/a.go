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

	if "A" <= s[0] && s[0] <= "Z" {
		fmt.Fprintf(writer, "%v\n", "A")
	} else {
		fmt.Fprintf(writer, "%v\n", "a")
	}
}