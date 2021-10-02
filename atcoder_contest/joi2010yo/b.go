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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}

	cur := 0
	for i := 0; i < m; i++ {
		cur += b[i]
		if cur >= n-1 {
			fmt.Fprintf(writer, "%v\n", i+1)
			return
		}
		cur += a[cur]
		if cur >= n-1 {
			fmt.Fprintf(writer, "%v\n", i+1)
			return
		}
	}

}
