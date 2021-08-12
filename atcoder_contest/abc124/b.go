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
	var h = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &h[i])
	}

	ans := 1
	for i := 1; i < n; i++ {
		ok := true
		for j := 0; j < i; j++ {
			// fmt.Fprintf(writer, "%v %v\n", i, j)
			if h[i] < h[j] {
				ok = false
				break
			}
		}
		if ok {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
