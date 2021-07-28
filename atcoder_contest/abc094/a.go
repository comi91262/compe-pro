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
	var x int
	fmt.Fscan(reader, &x)

	if a <= x && x <= a+b {
		fmt.Fprintf(writer, "%v\n", "YES")
	} else {
		fmt.Fprintf(writer, "%v\n", "NO")
	}
}
