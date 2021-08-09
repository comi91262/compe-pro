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
	var k int
	fmt.Fscan(reader, &k)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	l := b / k * k

	if a <= l {
		fmt.Fprintf(writer, "%v\n", "OK")
	} else {
		fmt.Fprintf(writer, "%v\n", "NG")
	}
}
