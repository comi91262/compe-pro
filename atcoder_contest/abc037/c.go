package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var b = make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = a[i]
	}
	for i := 0; i < n-1; i++ {
		b[i+1] += b[i]
	}

	ans := 0
	for i := 0; i < n-k+1; i++ {
		ans += b[i+k-1]

		if i != 0 {
			ans -= b[i-1]
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
