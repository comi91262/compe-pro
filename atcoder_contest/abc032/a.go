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

func lcm(a, b int) int {
	d := gcd(a, b)
	return a / d * b
}

func main() {
	defer writer.Flush()

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var n int
	fmt.Fscan(reader, &n)

	l := lcm(a, b)

	if n%l == 0 {
		fmt.Fprintf(writer, "%v\n", n)
	} else {
		fmt.Fprintf(writer, "%v\n", (n/l+1)*l)
	}
}
