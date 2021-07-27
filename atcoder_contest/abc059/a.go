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

	diff := "A"[0] - "a"[0]

	fmt.Fprintf(writer, "%v%v%v\n", string(a[0]+diff), string(b[0]+diff), string(c[0]+diff))
}

