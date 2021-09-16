package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

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

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}
func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	var l int
	fmt.Fscan(reader, &l)

	var p int
	fmt.Fscan(reader, &p)
	var q int
	fmt.Fscan(reader, &q)
	var r int
	fmt.Fscan(reader, &r)

	a := []int{n, m, l}
	b := []int{p, q, r}

	ans := 0
	for i := 0; i < 3*2; i++ {
		ans = max(ans, (a[0]/b[0])*(a[1]/b[1])*(a[2]/b[2]))
		nextPermutation(b)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
