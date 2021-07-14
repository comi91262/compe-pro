package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func lowerBound(value, boader int) bool {
	return boader <= value
}

func binarySearch(a []int, boader int, criteria func(value, boader int) bool) int {
	abs := func(x int) int {
		if x >= 0 {
			return x
		} else {
			return x * -1
		}
	}

	ng := -1
	ok := len(a)

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if criteria(a[mid], boader) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func unique(a []int) []int {
	r := make([]int, 0)
	m := make(map[int]bool, 0)

	for _, e := range a {
		if !m[e] {
			m[e] = true
			r = append(r, e)
		}
	}

	return r
}

func compression(a []int) []int {
	var n = len(a)
	var p = make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = a[i]
	}

	sort.Ints(p)
	p = unique(p)

	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = binarySearch(p, a[i], lowerBound)
	}

	return r
}

func sweepBitMatrix(a [][]int) [][]int {
	if row := len(a); row == 0 {
		return nil
	}

	row := len(a)
	col := len(a[0])

	cnt := make([]int, row)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if a[i][j] == 0 {
				cnt[i]++
				continue
			}
			break
		}

		if cnt[i] == col {
			continue
		}

		for j := 0; j < row; j++ {
			if i == j || a[j][cnt[i]] == 0 {
				continue
			}

			for k := 0; k < col; k++ {
				a[j][k] ^= a[i][k]
			}
		}
	}

	cnt = compression(cnt)
	c := make([][]int, row)

	for i := 0; i < row; i++ {
		c[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		c[cnt[i]] = a[i]
	}

	return c
}

func isZero(a []int) bool {
	for _, e := range a {
		if e != 0 {
			return false
		}
	}
	return true
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(reader, &t)

		for j := 0; j < t; j++ {
			var p int
			fmt.Fscan(reader, &p)
			a[i][p-1] = 1
		}
	}

	s := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &s[i])
	}

	a = sweepBitMatrix(a)

	//	for i := 0; i < 50; i++ {
	//		for j := 0; j < 50; j++ {
	//			fmt.Fprintf(writer, "%v ", a[i][j])
	//		}
	//		fmt.Fprintf(writer, "\n")
	//	}

	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < m; j++ {
			if s[j] == 1 {
				break
			}
			cnt++
		}

		if cnt < m && a[i][cnt] == 1 {
			for j := cnt; j < m; j++ {
				s[j] ^= a[i][j]
			}
		}
	}

	if isZero(s) {
		ans := 1
		for i := 0; i < n; i++ {
			if isZero(a[i]) {
				ans = ans * 2 % 998244353
			}
		}
		fmt.Fprintf(writer, "%v\n", ans)
	} else {
		fmt.Fprintf(writer, "%v\n", 0)
	}

}
