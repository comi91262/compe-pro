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

	for i := 0; i < len(s); i++ {
		fmt.Fprintf(writer, "%v", "x")

	}
	fmt.Fprintf(writer, "\n")
}
