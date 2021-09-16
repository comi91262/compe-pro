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
	var m int
	fmt.Fscan(reader, &m)
	var t = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &t[i])
	}

	ma := map[int]int{}
	for i := 0; i < n; i++ {
		ma[d[i]]++
	}

	for i := 0; i < m; i++ {
		if ma[t[i]] == 0 {
			fmt.Fprintf(writer, "%v\n", "NO")
			return
		}
		ma[t[i]]--
	}

	fmt.Fprintf(writer, "%v\n", "YES")
}
