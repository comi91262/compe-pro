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

	var p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}

	for i := 0; i < n; i++ {
		if i+1 != p[i] {
			p[p[i]-1], p[i] = p[i], p[p[i]-1]
			break
		}
	}

	for i := 0; i < n-1; i++ {
		if i+1 != p[i] {
			fmt.Fprintf(writer, "%v\n", "NO")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "YES")

}
