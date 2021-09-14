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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	total := 0
	for i := 0; i < n; i++ {
		if total <= k && k <= total+a[i] {
			fmt.Fprintf(writer, "%v\n", i+1)
			return
		}
		total += a[i]
	}
}
