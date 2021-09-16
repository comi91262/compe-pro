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
	var r int
	fmt.Fscan(reader, &r)
	var g int
	fmt.Fscan(reader, &g)
	var b int
	fmt.Fscan(reader, &b)
	var n int
	fmt.Fscan(reader, &n)

	ans := 0
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			k := (n - i*r - j*g) / b
			if 0 <= k && i*r+j*g+k*b == n {
				ans++
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
