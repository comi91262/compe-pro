package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	ans := 0
	for {
		r := int(math.Sqrt(float64(n)))
		if r*r == n {
			break
		}

		n++
		ans++
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
