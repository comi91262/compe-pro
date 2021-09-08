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
	var s = make([]string, n)
	var t = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
		fmt.Fscan(reader, &t[i])
	}
	var x string
	fmt.Fscan(reader, &x)

	st := 0
	for i := 0; i < n; i++ {
		if s[i] == x {
			st = i + 1
			break
		}
	}
	ans := 0
	for i := st; i < n; i++ {
		ans += t[i]
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
