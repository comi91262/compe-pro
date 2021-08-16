package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func calc(x int) int {
	b := 0
	switch x % 4 {
	case 0:
		b ^= x
	case 1:
		b ^= x ^ (x - 1)
	case 2:
		b ^= x ^ (x - 1) ^ (x - 2)
	case 3:
		b = 0
	}
	return b
}

func main() {
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	if a == 0 {
		fmt.Fprintf(writer, "%v\n", calc(b))
	} else {
		fmt.Fprintf(writer, "%v\n", calc(b)^calc(a-1))
	}
}
