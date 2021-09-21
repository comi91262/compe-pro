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
	var a = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = i + 1
	}

	hand := 0

	for i := 0; i < m; i++ {
		idx := -1
		for j := 0; j < len(ans); j++ {
			if a[i] == ans[j] {
				idx = j
				break
			}
		}
		if idx == -1 {
			continue
		}
		hand, ans[idx] = ans[idx], hand
	}
	for i := range ans {
		fmt.Fprintf(writer, "%v\n", ans[i])
	}
}
