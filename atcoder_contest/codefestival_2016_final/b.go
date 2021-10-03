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

	s := 0
	m := 1
	for s+m < n {
		s += m
		m++
	}
	s += m

	ans := []int{}
	for i := 1; i <= m; i++ {
		if i == s-n {
			continue
		}
		ans = append(ans, i)
	}

	for i := range ans {
		fmt.Fprintf(writer, "%v\n", ans[i])
	}

}
