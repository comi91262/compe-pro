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

	a := make([]int, 50)
	a[0] = 1
	a[1] = 1
	a[2] = 1
	a[3] = 1
	for i := 4; i < 50; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	ans := 0
	for i := 0; i <= n; i++ {
		ans += a[i]
	}

	if n == 0 || n == 1 {
		fmt.Fprintf(writer, "%v\n", 1)
	} else {
		fmt.Fprintf(writer, "%v\n", ans-1)
	}
}

