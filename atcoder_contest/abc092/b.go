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

	var d int
	fmt.Fscan(reader, &d)
	var x int
	fmt.Fscan(reader, &x)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		r := 0
		for j := 1; j <= d; j += a[i] {
			r += 1
		}
		ans += r
	}
	fmt.Fprintf(writer, "%v\n", ans+x)

}
