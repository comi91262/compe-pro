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

	var a int
	fmt.Fscan(reader, &a)
	var v int
	fmt.Fscan(reader, &v)
	var b int
	fmt.Fscan(reader, &b)
	var w int
	fmt.Fscan(reader, &w)
	var t int
	fmt.Fscan(reader, &t)

	if a < b {
		if b+w*t <= a+v*t {
			fmt.Fprintf(writer, "%v\n", "YES")
			return
		}
	} else {
		if a-v*t <= b-w*t {
			fmt.Fprintf(writer, "%v\n", "YES")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "NO")
}
