package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func min(x, y, z int) int {
	var m = int(math.Min(float64(x), float64(y)))
	return int(math.Min(float64(m), float64(z)))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)

	var n, m, l = gcd(a, b), gcd(b, c), gcd(c, a)
	var u = min(n, m, l)

	fmt.Fprintf(writer, "%d\n", a/u-1+b/u-1+c/u-1)
}
