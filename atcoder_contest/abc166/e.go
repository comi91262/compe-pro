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
	{
		b := make([]int, n)
		c := make([]int, n)

		for i := 0; i < n; i++ {
			b[i] = i + 1 + a[i]
			c[i] = i + 1 - a[i]
		}

		m := map[int]int{}
		m[b[0]]++
		for i := 1; i < n; i++ {
			if m[c[i]] > 0 {
				ans += m[c[i]]
			}
			m[b[i]]++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
