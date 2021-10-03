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
	mp := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		mp[a[i]]++
	}
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
		mp[b[i]]++
	}

	fmt.Fprintf(writer, "%v\n", float64(n+m-len(mp))/float64(len(mp)))
}
