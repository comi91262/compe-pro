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

	ans := []int{}
	for i := 1; i <= 127; i++ {
		if i%3 == a && i%5 == b && i%7 == c {
			ans = append(ans, i)
		}
	}
	for i := range ans {
		fmt.Fprintf(writer, "%v\n", ans[i])
	}
}
