package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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

	m := map[int]int{}
	for i := 2; i <= n; i++ {
		l := primeFactor(i)
		for k, v := range l {
			m[k] += v
		}
	}

	a := []int{2, 4, 14, 24, 74}
	m2 := map[int]int{}
	for _, v := range m {
		for i := range a {
			if v >= a[i] {
				m2[a[i]]++
			}
		}
	}

	ans := 0
	ans += m2[4] * (m2[4] - 1) / 2 * (m2[2] - 2)
	ans += m2[74]
	ans += m2[24] * (m2[2] - 1)
	ans += m2[14] * (m2[4] - 1)

	fmt.Fprintf(writer, "%v\n", ans)
}
