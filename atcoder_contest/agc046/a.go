package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func main() {
	defer writer.Flush()
	var x int
	fmt.Fscan(reader, &x)
	fmt.Fprintf(writer, "%v\n", 360/gcd(360, x))
}

