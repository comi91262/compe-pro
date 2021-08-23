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
	var p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}

	ans := 0
	for i := 0; i < n-2; i++ {
		if p[i] <= p[i+1] && p[i+1] <= p[i+2] {
			ans++
		}
		if p[i] >= p[i+1] && p[i+1] >= p[i+2] {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
