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

	var n string
	fmt.Fscan(reader, &n)

	for strings.HasSuffix(n, "0") {
		n = n[:len(n)-1]
	}

	for i := 0; i < len(n)/2; i++ {
		if n[i] != n[len(n)-i-1] {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "Yes")
}
