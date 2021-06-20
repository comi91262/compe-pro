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

	sum := 0
	ans := 1
	for i := 1; ; i++ {
		sum += i
		if sum >= n {
			ans = i
			break
		}
	}

	fmt.Fprintf(writer, "%d\n", ans)
}
