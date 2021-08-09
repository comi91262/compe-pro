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

	var d int
	fmt.Fscan(reader, &d)
	var s int
	fmt.Fscan(reader, &s)
	var t int
	fmt.Fscan(reader, &t)

	if d <= s*t {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}

}
