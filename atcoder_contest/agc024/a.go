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
	var c int
	fmt.Fscan(reader, &c)
	var k int
	fmt.Fscan(reader, &k)

	if a == b {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}
	if k%2 == 0 {
		fmt.Fprintf(writer, "%v\n", a-b)
	} else {
		fmt.Fprintf(writer, "%v\n", b-a)
	}
}
