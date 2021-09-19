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
	var k int
	fmt.Fscan(reader, &k)

	ans := 0
	for a := 1; a <= k; a++ {
		for b := 1; a*b <= k; b++ {
			ans += k / (a * b)
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
