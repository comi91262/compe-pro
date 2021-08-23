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

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 0
	cnt := 0
	pre := 0
	for i := 0; i < n; i++ {
		if a[i] <= pre {
			ans += cnt * (cnt - 1) / 2
			cnt = 0
		}
		ans++
		cnt++

		pre = a[i]
	}
	if cnt != 0 {
		ans += cnt * (cnt - 1) / 2
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
