package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	if n%2 == 1 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}
	cnt := []int{}
	base := 10
	for base <= n {
		cnt = append(cnt, n/base)
		base *= 5
	}
	fmt.Fprintf(writer, "%v\n", sum(cnt))
}
