package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// スライス a の i番目からj番目を反転させる関数
func reverse(a []int, i, j int) {
	for k := 0; k < (j-i+1)/2; k++ {
		swap(a, i+k, j-k)
	}
}

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

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var x = make([]float64, n)
	var y = make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		fmt.Fscan(reader, &y[i])
	}

	a := []int{}
	for i := 0; i < n; i++ {
		a = append(a, i)
	}

	sum := 0.0
	for i := 0; i < fact(n); i++ {
		for j := 0; j < len(a)-1; j++ {
			s := a[j]
			t := a[j+1]
			dx := x[s] - x[t]
			dy := y[s] - y[t]
			sum += math.Sqrt(dx*dx + dy*dy)
		}
		nextPermutation(a)
	}
	fmt.Fprintf(writer, "%v\n", sum/float64(fact(n)))

}
