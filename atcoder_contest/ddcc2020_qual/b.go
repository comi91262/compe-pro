package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func add(arg ...int) (sum int) {
	for i := range arg {
		sum += arg[i]
	}
	return
}
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
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
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	s := 0
	t := add(a...)
	ans := 1 << 60
	for i := 0; i < n; i++ {
		s += a[i]
		t -= a[i]
		ans = min(abs(s-t), ans)
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
