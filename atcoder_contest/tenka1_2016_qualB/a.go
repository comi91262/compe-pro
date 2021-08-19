package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func f(x int) int {
	return (x*x + 4) / 8
}

func main() {
	defer writer.Flush()
	fmt.Fprintf(writer, "%v\n", f(f(f(20))))
}
