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
	var n int
	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(reader, &a)
		var b int
		fmt.Fscan(reader, &b)

		fmt.Fprintf(writer, "%v\n", gcd(a, b))
	}

}
