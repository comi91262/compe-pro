package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	sum := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			sum += i*100 + j
		}
	}

	fmt.Fprintf(writer, "%d\n", sum)
}
