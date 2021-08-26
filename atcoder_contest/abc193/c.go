package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
}
func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	ans := 0
	m := map[int]int{}
	for a := 2; a*a <= n; a++ {
		for b := 2; ; b++ {
			if n < pow(a, b) || m[pow(a, b)] > 0 {
				break
			}
			m[pow(a, b)]++
			ans++
		}

	}

	fmt.Fprintf(writer, "%v\n", n-ans)
}
