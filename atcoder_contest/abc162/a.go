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

	for i := range n {
		if n[i] == "7"[0] {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}

	}

	fmt.Fprintf(writer, "%v\n", "No")
}
