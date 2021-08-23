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
	fmt.Fscan(reader, &ss)
	var t = strings.Split(ss, "")

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v%v", s[i], t[i])
	}
	fmt.Fprintf(writer, "\n")

}
