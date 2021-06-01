package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}

	diff := 0
	for i := 0; i < n; i++ {
		diff += abs(a[i] - b[i])
	}

	if diff > k {
		fmt.Fprintf(writer, "No\n")
	} else if diff == k {
		fmt.Fprintf(writer, "Yes\n")
	} else {
		if (k+diff)%2 == 0 {
			fmt.Fprintf(writer, "Yes\n")
		} else {
			fmt.Fprintf(writer, "No\n")
		}
	}
}
