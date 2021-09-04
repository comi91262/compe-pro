package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	mx := 0
	for i := 0; i < n; i++ {
		mx = max(mx, a[i])
	}
	t := make([]bool, mx+1)

	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]]++
	}
	sort.Ints(a)

	for i := 0; i < len(a); i++ {
		if m[a[i]] > 1 {
			for j := a[i]; j <= mx; j += a[i] {
				t[j] = true
			}
			continue
		}
		for j := a[i] + a[i]; j <= mx; j += a[i] {
			t[j] = true
		}
	}

	cnt := 0
	for i := 0; i < len(a); i++ {
		if !t[a[i]] {
			cnt++
		}
	}
	fmt.Fprintf(writer, "%v\n", cnt)
}
