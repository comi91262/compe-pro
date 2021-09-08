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
	var m = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &m[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		if m[i] >= 80 {
			continue
		}
		ans += 80 - m[i]
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
