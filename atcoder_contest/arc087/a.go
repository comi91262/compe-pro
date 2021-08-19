package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
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
	for k, v := range m {
		if k > v {
			ans += v
		} else if k < v {
			ans += abs(k - v)
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
