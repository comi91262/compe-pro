package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func lis(a []int) (lis []int) {
	lis = make([]int, len(a))

	b := make([]int, 0)
	for i := 0; i < len(a); i++ {
		cnt := sort.Search(len(b), func(j int) bool { return a[i] < b[j] })

		if cnt == len(b) {
			b = append(b, a[i])
		} else {
			b[cnt] = a[i]
		}
		lis[i] = cnt + 1
	}

	return
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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 0
	for _, v := range lis(a) {
		ans = max(ans, v)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
