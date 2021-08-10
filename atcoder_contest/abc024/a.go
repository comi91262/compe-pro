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
	var k int
	fmt.Fscan(reader, &k)
	var s int
	fmt.Fscan(reader, &s)
	var t int
	fmt.Fscan(reader, &t)

	ans := 0
	if s+t >= k {
		ans = s*(a-c) + t*(b-c)

	} else {
		ans = s*a + t*b
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
