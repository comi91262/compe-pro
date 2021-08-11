package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func divisor(n int) []int {
	var r []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			r = append(r, i)
			if i*i != n {
				r = append(r, n/i)
			}
		}
	}
	return r
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	ans := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			continue
		}
		if len(divisor(i)) == 8 {
			ans++
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
