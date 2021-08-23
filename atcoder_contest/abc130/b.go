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
	var x int
	fmt.Fscan(reader, &x)
	var l = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &l[i])
	}

	ans := 0
	d := 0

	for i := 0; i <= n; i++ {
		if i != 0 {
			d = d + l[i-1]
		}
		if d <= x {
			ans++
		}

	}
	fmt.Fprintf(writer, "%v\n", ans)
}
