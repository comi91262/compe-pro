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

	var n string
	fmt.Fscan(reader, &n)

	if n[0] == n[1] && n[1] == n[2] {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else if n[1] == n[2] && n[2] == n[3] {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
