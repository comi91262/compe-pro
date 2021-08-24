package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}

	ans := 0
	for i := 0; i < n; i++ {

		if a[i] <= b[i] {
			ans += a[i]
			b[i] -= a[i]
			a[i] = 0
		} else {
			ans += b[i]
			a[i] -= b[i]
			b[i] = 0
			continue
		}

		if a[i+1] <= b[i] {
			ans += a[i+1]
			b[i] -= a[i+1]
			a[i+1] = 0
		} else {
			ans += b[i]
			a[i+1] -= b[i]
			b[i] = 0
			continue
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
