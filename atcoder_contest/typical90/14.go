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
func min(x, y, z int) int {
	var m = int(math.Min(float64(x), float64(y)))
	return int(math.Min(float64(m), float64(z)))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
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
