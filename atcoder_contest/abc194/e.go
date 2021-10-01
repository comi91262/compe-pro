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

func mex(m map[int]int) int {
	for i := 0; ; i++ {
		if m[i] == 0 {
			return i
		}
	}
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

	mp := map[int]int{}
	for i := 0; i < m; i++ {
		mp[a[i]]++
	}

	ans := mex(mp)
	for i := 1; i < n-m+1; i++ {
		mp[a[i-1]]--
		mp[a[i+m-1]]++
		if mp[a[i-1]] == 0 {
			ans = min(a[i-1], ans)
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
