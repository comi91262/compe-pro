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

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		if a[i] > 10 {
			sum += a[i] - 10
		}
	}

	fmt.Fprintf(writer, "%d\n", sum)
}
