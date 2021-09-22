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

	if ((n-1)/20)%2 == 0 {
		fmt.Fprintf(writer, "%v\n", (n-1)%20+1)
	} else {
		fmt.Fprintf(writer, "%v\n", 21-((n-1)%20+1))
	}
}
