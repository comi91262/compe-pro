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
	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]]++
	}

	ans := 0
	for k1, v1 := range m {
		for k2, v2 := range m {
			if k1 < k2 {
				continue
			}
			if k1 == k2 {
				continue
			}

			ans += (k1 - k2) * (k1 - k2) * v1 * v2
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}

