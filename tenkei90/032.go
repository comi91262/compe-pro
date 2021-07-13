package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func reverse(a []int, i, j int) {
	for k := 0; k < (j-i+1)/2; k++ {
		swap(a, i+k, j-k)
	}
}

func factorial(n int) int {
	var prod = 1
	for i := 1; i <= n; i++ {
		prod *= i
	}
	return prod
}

// Narayana Pandita’s algorithm
func nextPermutation(a []int) {
	var lastIndex = len(a) - 1

	if lastIndex < 1 {
		return
	}

	var i = lastIndex - 1
	for i >= 0 && a[i] >= a[i+1] {
		i -= 1
	}
	if i < 0 {
		reverse(a, 0, lastIndex)
		return
	}

	var j = lastIndex
	for j > i+1 && a[j] <= a[i] {
		j -= 1
	}

	swap(a, i, j)
	reverse(a, i+1, lastIndex)
}

func check(c, x, y []int, n, m int) bool {
	for j := 0; j < m; j++ {
		for k := 0; k < n; k++ {
			if x[j] == c[k] {
				if k > 0 && c[k-1] == y[j] {
					return false
				}
				if k+1 < n && c[k+1] == y[j] {
					return false
				}
			}

		}
	}

	return true

}

const MIN = 10001

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([][]int, 0, n)
	for i := 0; i < n; i++ {
		var b = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &b[j])
		}
		a = append(a, b)
	}

	var m int
	fmt.Fscan(reader, &m)

	var x = make([]int, m)
	var y = make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}

	var c = make([]int, 0, n)
	for i := 0; i < n; i++ {
		c = append(c, i+1)
	}

	mini := MIN
	fact := factorial(n)

	for i := 0; i < fact; i++ {
		if check(c, x, y, n, m) {
			sum := 0
			for i := 0; i < n; i++ {
				sum += a[c[i]-1][i]
			}
			mini = min(mini, sum)
		}
		nextPermutation(c)
	}

	if mini == MIN {
		fmt.Fprintf(writer, "%d\n", -1)
	} else {
		fmt.Fprintf(writer, "%d\n", mini)
	}
}
