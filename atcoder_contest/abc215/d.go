package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
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

func primeFactor(n int) map[int]int {
	var m = map[int]int{}

	for i := 2; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n != 1 {
		m[n]++
	}
	return m
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ng := make([]bool, 1_00001)
	ng[0] = true
	for i := 0; i < n; i++ {
		for k := range primeFactor(a[i]) {
			ng[k] = true
		}
	}

	for i := 2; i <= m; i++ {
		if ng[i] {
			for y := i + i; y <= m; y += i {
				ng[y] = true
			}
		}
	}
	ans := []int{}

	for i := 1; i <= m; i++ {
		if !ng[i] {
			ans = append(ans, i)
		}
	}

	fmt.Fprintf(writer, "%v\n", len(ans))
	for i := range ans {
		fmt.Fprintf(writer, "%v\n", ans[i])
	}

}
