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

	var n float64
	fmt.Fscan(reader, &n)
	var k float64
	fmt.Fscan(reader, &k)

	ans := 0.0

	for i := 0; i < int(n); i++ {
		p := 1 / n
		cnt := i + 1
		for cnt < int(k) {
			p *= 0.5
			cnt *= 2
		}
		ans += p
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
