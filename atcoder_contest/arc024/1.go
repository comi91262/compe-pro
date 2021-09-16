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

func main() {
	defer writer.Flush()
	var l, r int
	fmt.Fscan(reader, &l, &r)
	var a = make([]int, l)
	s := map[int]int{}
	for i := 0; i < l; i++ {
		fmt.Fscan(reader, &a[i])
		s[a[i]]++
	}
	var b = make([]int, r)
	t := map[int]int{}
	for i := 0; i < r; i++ {
		fmt.Fscan(reader, &b[i])
		t[b[i]]++
	}

	ans := 0
	for k, v := range s {
		if t[k] > 0 {
			ans += min(t[k], v)
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
