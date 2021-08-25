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
	var k int
	fmt.Fscan(reader, &k)
	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var b = make([]int, n)
	for i := 0; i < n-1; i++ {
		b[i] = a[i+1] - a[i]
	}
	b[n-1] = k - a[n-1] + a[0]

	mx := 0
	for i := 0; i < n; i++ {
		mx = max(mx, b[i])
	}
	fmt.Fprintf(writer, "%v\n", k-mx)

}
