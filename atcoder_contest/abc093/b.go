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

func main() {
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var k int
	fmt.Fscan(reader, &k)

	r := []int{}
	for i := a; i <= min(a+k-1, b); i++ {
		r = append(r, i)
	}

	for i := max(b-k+1, a+k); i <= b; i++ {
		r = append(r, i)
	}

	// m := map[int]int{}
	// for i := range r {
	// 	m[r[i]]++
	// }

	// ans := []int{}
	// for k := range m {
	// 	ans = append(ans, k)
	// }

	// sort.Ints(ans)

	for i := range r {
		fmt.Fprintf(writer, "%v\n", r[i])
	}
}
