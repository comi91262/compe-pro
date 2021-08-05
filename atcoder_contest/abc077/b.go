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

	for i := 0; ; i++ {
		x := int(math.Sqrt(float64(n - i)))
		if x*x == n-i {
			fmt.Fprintf(writer, "%v\n", n-i)
			return
		}
	}
}
