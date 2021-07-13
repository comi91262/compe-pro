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

	var n, q int
	fmt.Fscan(reader, &n, &q)

	var a = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var b = make([]int, n+1)
	for i := 1; i < n; i++ {
		b[i] = a[i+1] - a[i]
	}

	sum := 0
	for i := 1; i < n; i++ {
		sum += abs(b[i])
	}

	var l = make([]int, q+1)
	var r = make([]int, q+1)
	var v = make([]int, q+1)
	for i := 1; i < q+1; i++ {
		fmt.Fscan(reader, &l[i], &r[i], &v[i])
	}

	for i := 1; i < q+1; i++ {
		mae := abs(b[l[i]-1]) + abs(b[r[i]])
		if l[i] >= 2 {
			b[l[i]-1] += v[i]
		}
		if r[i] <= n-1 {
			b[r[i]] -= v[i]
		}
		ato := abs(b[l[i]-1]) + abs(b[r[i]])
		sum += (ato - mae)
		fmt.Fprintf(writer, "%d\n", sum)
	}

}
