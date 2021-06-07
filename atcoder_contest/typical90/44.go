package main

import (
	"bufio"
	"fmt"
	"os"
)

var ans struct {
	a     []int
	shift int
}

func mod(x, y int) int {
	x = x % y
	if x >= 0 {
		return x
	}
	if y < 0 {
		return x - y
	}
	return x + y
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	ans.a = a
	ans.shift = 0

	var t = make([]int, q)
	var x = make([]int, q)
	var y = make([]int, q)

	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &t[i], &x[i], &y[i])
	}

	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			var xIdx = mod(x[i]-1-ans.shift, n)
			var yIdx = mod(y[i]-1-ans.shift, n)
			ans.a[xIdx], ans.a[yIdx] = ans.a[yIdx], ans.a[xIdx]
		case 2:
			ans.shift++
		case 3:
			var xIdx = mod(x[i]-1-ans.shift, n)
			fmt.Fprintf(writer, "%d\n", ans.a[xIdx])
		}
	}
}
