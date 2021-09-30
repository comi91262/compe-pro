package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	defer writer.Flush()

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	ans := 0
	for i := 1; i < 200001; i++ {
		if b/i-(a-1)/i >= 2 {
			ans = max(ans, i)
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
