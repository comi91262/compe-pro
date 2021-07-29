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

	var s string
	fmt.Fscan(reader, &s)

	mn := 1 << 60
	mx := 0
	for i := range s {
		switch s[i] {
		case "A"[0]:
			mn = min(mn, i)
		case "Z"[0]:
			mx = max(mx, i)
		}
	}

	fmt.Fprintf(writer, "%v\n", mx-mn+1)

}
