package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	ans := 0
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Fscan(reader, &tmp)

		ans += 2 * min(tmp, k-tmp)
	}

	fmt.Fprintf(writer, "%v\n", ans)
}


