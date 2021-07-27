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

	if n[0] == n[len(n)-1] {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
