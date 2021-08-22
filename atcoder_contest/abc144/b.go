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

	var n int
	fmt.Fscan(reader, &n)

	for i := 1; i < 10; i++ {
		if n/i*i == n && 1 <= n/i && n/i < 10 {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "No")
}
