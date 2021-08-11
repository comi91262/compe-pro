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
	var t int
	fmt.Fscan(reader, &t)

	var c = make([]int, n)
	var d = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
		fmt.Fscan(reader, &d[i])
	}

	ans := 1 << 60
	for i := 0; i < n; i++ {
		if d[i] <= t {
			ans = min(ans, c[i])
		}
	}

	if ans == 1<<60 {
		fmt.Fprintf(writer, "%v\n", "TLE")
	} else {
		fmt.Fprintf(writer, "%v\n", ans)
	}

}
