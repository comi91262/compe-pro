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
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]]++
	}

	if len(m)%2 == 0 {
		fmt.Fprintf(writer, "%v\n", len(m)-1)
	} else {
		fmt.Fprintf(writer, "%v\n", len(m))
	}
}
