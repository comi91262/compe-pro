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
	var k int
	fmt.Fscan(reader, &k)

	ans := 0
	for i := 1; i <= k; i++ {
		for j := 1; j <= k; j++ {
			for l := 1; l <= k; l++ {
				ans += gcd(i, gcd(j, l))
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
