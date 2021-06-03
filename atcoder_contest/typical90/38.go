package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b uint64) uint64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b uint64
	fmt.Fscan(reader, &a, &b)

	var g = gcd(a, b)

	if b/g > 1_000_000_000_000_000_000/a {
		fmt.Fprintf(writer, "Large")
	} else {
		fmt.Fprintf(writer, "%d\n", a/g*b)
	}

}
