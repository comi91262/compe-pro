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
	var a = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &a[i])
	}

	m := map[int]int{}
	for i := 0; i < n-1; i++ {
		m[a[i]]++
	}
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", m[i+1])
	}
}
