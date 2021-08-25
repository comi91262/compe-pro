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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		step := -(a[i+1] - a[i])
		if step <= 0 {
			continue
		}
		ans += step
		a[i+1] += step
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
