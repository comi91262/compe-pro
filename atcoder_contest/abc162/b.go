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
	ans := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			continue
		}
		if i%5 == 0 {
			continue
		}
		if i%15 == 0 {
			continue
		}

		ans += i
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
