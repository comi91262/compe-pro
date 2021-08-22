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
	var d = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if j == i {
				continue
			}
			ans += d[i] * d[j]
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
