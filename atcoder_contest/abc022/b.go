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

	m := map[int]int{}
	ans := 0
	for i := 0; i < n; i++ {
		if m[a[i]] > 0 {
			ans++
		}
		m[a[i]]++
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
