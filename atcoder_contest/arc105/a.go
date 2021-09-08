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
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)
	var d int
	fmt.Fscan(reader, &d)

	e := []int{a, b, c, d}
	n := len(e)

	for i := 0; i < 1<<n; i++ {
		sum := 0
		rest := 0
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				sum += e[j]
			} else {
				rest += e[j]
			}
		}
		if sum == rest {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "No")

}
