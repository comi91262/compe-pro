package main

import "fmt"

func step2(n, l, k, m int, a []int) bool {
	cnt := 0
	pre := 0
	for i := 0; i < n; i++ {
		if a[i]-pre >= m && l-a[i] >= m {
			cnt++
			pre = a[i]
		}
	}
	return k <= cnt
}

func main() {
	var n, l, k int
	fmt.Scan(&n, &l, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	m := l / 2
	upper := l
	lower := 0
	for {
		if step2(n, l, k, m, a) {
			if m == (upper-m)/2+m {
				break
			}
			lower = m
			m = (upper-m)/2 + m
		} else {
			if m == m/2 {
				break
			}
			upper = m
			m = (m - lower) / 2
		}
	}
	fmt.Println(m)
}