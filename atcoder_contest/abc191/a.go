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
	var v int
	fmt.Fscan(reader, &v)
	var t int
	fmt.Fscan(reader, &t)
	var s int
	fmt.Fscan(reader, &s)
	var d int
	fmt.Fscan(reader, &d)

	if v*t <= d && d <= v*s {
		fmt.Fprintf(writer, "%v\n", "No")
	} else {
		fmt.Fprintf(writer, "%v\n", "Yes")
	}

}
