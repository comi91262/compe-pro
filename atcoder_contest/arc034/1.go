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
	ans := 0.0
	for i := 0; i < n; i++ {
		var a, b, c, d, e float64
		fmt.Fscan(reader, &a, &b, &c, &d, &e)
		ans = math.Max(ans, a+b+c+d+e*110.0/900.0)
	}

    fmt.Fprintf(writer, "%v\n", ans)
}
