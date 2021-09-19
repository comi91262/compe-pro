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
		a[i]--
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Fprintf(writer, " ")
		}
		b := make([]int, n)
		copy(b, a)

		cnt := 0
		x := i
		once := true
		for once || x != i {
			x, b[x] = b[x], x
			once = false
			cnt++
		}
		fmt.Fprintf(writer, "%v", cnt)
	}
	fmt.Fprintf(writer, "\n")
}
