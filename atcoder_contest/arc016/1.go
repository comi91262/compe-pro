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
	var m int
	fmt.Fscan(reader, &m)

	ans := []int{}
	for i := 0; i < n; i++ {
		ans = append(ans, i+1)
	}
	for i := range ans {
		if ans[i] != m {
			fmt.Fprintf(writer, "%v\n", ans[i])
			return
		}
	}
}
