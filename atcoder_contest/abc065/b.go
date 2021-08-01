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

	l := 0
	i := 0
	for i <= 500000 {
		l = a[l] - 1
		i++
		if l == 1 {
			fmt.Fprintf(writer, "%v\n", i)
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", -1)

}
