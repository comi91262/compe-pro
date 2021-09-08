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

	if strings.Contains("OPKL", s) {
		fmt.Fprintf(writer, "%v\n", "Right")
	} else {
		fmt.Fprintf(writer, "%v\n", "Left")
	}
}
