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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var b = make([]int, n)
	b[0] = a[0]
	for i := 1; i < n; i++ {
		b[i] = min(b[i-1], a[i])
	}

	ans := 1
	for i := 1; i < n; i++ {
		if b[i-1] >= a[i] {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
