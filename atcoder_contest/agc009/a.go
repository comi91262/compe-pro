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
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	cnt := 0
	for i := 0; i < n; i++ {
		inc := a[n-1-i] + cnt
		nex := (inc+b[n-1-i]-1)/b[n-1-i]*b[n-1-i] - inc
		cnt += nex
	}
	fmt.Fprintf(writer, "%v\n", cnt)

}

