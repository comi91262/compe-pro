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
	var k int
	fmt.Fscan(reader, &k)

	ma := map[int]int{}
	ma[0]++
	ma[n*m]++

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			ma[i*j+(n-i)*(m-j)]++
		}
	}
	if ma[k] > 0 {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
