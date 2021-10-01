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
	var m int
	fmt.Fscan(reader, &m)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	mp := map[int]int{}
	for i := 0; i < n; i++ {
		mp[a[i]]++
	}
	ans := -1
	for point, cnt := range mp {
		if cnt > n/2 {
			ans = point
		}
	}
	if ans == -1 {
		fmt.Fprintf(writer, "%v\n", "?")
	} else {
		fmt.Fprintf(writer, "%v\n", ans)
	}

}
