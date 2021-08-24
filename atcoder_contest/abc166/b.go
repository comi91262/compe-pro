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
	var k int
	fmt.Fscan(reader, &k)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		var d int
		fmt.Fscan(reader, &d)

		a[i] = make([]int, d)
		for j := 0; j < d; j++ {
			fmt.Fscan(reader, &a[i][j])
			a[i][j]--
		}
	}

	var b = make([]int, n)
	for i := 0; i < n; i++ {
		for j := range a[i] {
			b[a[i][j]]++
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if b[i] == 0 {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
