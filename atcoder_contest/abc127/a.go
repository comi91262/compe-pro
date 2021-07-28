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
	var b int
	fmt.Fscan(reader, &b)

	switch {
	case a >= 13:
		fmt.Fprintf(writer, "%v\n", b)
	case 6 <= a && a <= 12:
		fmt.Fprintf(writer, "%v\n", b/2)
	case a <= 5:
		fmt.Fprintf(writer, "%v\n", 0)
	}

}
