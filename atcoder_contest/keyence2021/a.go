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
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}

	amax := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			amax[i] = a[i]
			continue
		}
		amax[i] = max(amax[i-1], a[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, b[i]*amax[i])
		fmt.Fprintf(writer, "%v\n", ans)
	}
}

