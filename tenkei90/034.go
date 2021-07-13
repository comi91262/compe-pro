package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 0
	cr := 0
	cnt := 0
	m := map[int]int{}
	for i := 0; i < n; i++ {
		for cr < n {
			if m[a[cr]] == 0 && cnt == k {
				break
			}
			if m[a[cr]] == 0 {
				cnt++
			}
			m[a[cr]] += 1
			cr++
		}
		ans = max(ans, cr-i)
		if m[a[i]] == 1 {
			cnt--
		}
		m[a[i]] -= 1
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
