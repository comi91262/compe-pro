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

	for i := 0; i < n; i++ {
		if a[i] == 0 {
			fmt.Fprintf(writer, "%v\n", 0)
			return
		}
	}

	ans := 1
	for i := 0; i < n; i++ {
		if ans > 1_000_000_000_000_000_000/a[i] {
			ans = -1
			break
		}
		ans *= a[i]
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
