package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	ans := []int{}
	for i := 0; i < 1<<n; i++ {
		r := 1
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				r *= 2
			} else {
				r += k
			}
		}
		ans = append(ans, r)
	}
	sort.Ints(ans)

	fmt.Fprintf(writer, "%v\n", ans[0])
}
