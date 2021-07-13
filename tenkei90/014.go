package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

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
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}

	sort.Ints(a)
	sort.Ints(b)

	sum := 0
	for i := 0; i < n; i++ {
		sum += abs(a[i] - b[i])

	}

	fmt.Fprintf(writer, "%d\n", sum)
}
