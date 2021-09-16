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
	var t = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
	}

	for i := 0; i < n-3+1; i++ {
		if k > t[i]+t[i+1]+t[i+2] {
			fmt.Fprintf(writer, "%v\n", i+1+2)
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", -1)

}
