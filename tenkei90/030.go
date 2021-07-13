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

	var a = make([]int, n+1)

	for i := 2; i < n+1; i++ {
		if a[i] != 0 {
			continue
		}

		for j := i; j < n+1; j += i {
			a[j]++
		}
	}
	sum := 0
	for i := 1; i < n+1; i++ {
		if a[i] >= k {
			sum++
		}
	}

	fmt.Fprintf(writer, "%d\n", sum)
}
