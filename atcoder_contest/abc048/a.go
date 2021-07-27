package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var a string
	fmt.Fscan(reader, &a)
	var b string
	fmt.Fscan(reader, &b)
	var c string
	fmt.Fscan(reader, &c)
	fmt.Fprintf(writer, "%s%s%s\n", string(a[0]), string(b[0]), string(c[0]))
}
