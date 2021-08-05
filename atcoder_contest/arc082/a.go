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
		m[a[i]-1]++
		m[a[i]+1]++
	}

	ans := 0
	for _, v := range m {
		ans = max(ans, v)
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}
