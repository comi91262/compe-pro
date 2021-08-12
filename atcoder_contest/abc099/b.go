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

	c := b - a
	sum := 0
	for i := 1; i < c; i++ {
		sum += i
	}
	fmt.Fprintf(writer, "%v\n", sum-a)

}