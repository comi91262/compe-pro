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
	var x int
	fmt.Fscan(reader, &x)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		if a[i]+a[i+1] > x {
			rest := a[i] + a[i+1] - x
			if a[i+1] >= rest {
				a[i+1] -= rest
			} else {
				a[i] = rest - a[i+1]
				a[i+1] = 0
			}
			ans += rest
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
