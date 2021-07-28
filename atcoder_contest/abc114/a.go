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

	var x int
	fmt.Fscan(reader, &x)

	switch x {
	case 3:
		fmt.Fprintf(writer, "%v\n", "YES")
	case 5:
		fmt.Fprintf(writer, "%v\n", "YES")
	case 7:
		fmt.Fprintf(writer, "%v\n", "YES")
	default:
		fmt.Fprintf(writer, "%v\n", "NO")
	}

}
