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
	var b int
	fmt.Fscan(reader, &b)

	fmt.Fprintf(writer, "%v\n", s[(b-1)%len(s)])
}
