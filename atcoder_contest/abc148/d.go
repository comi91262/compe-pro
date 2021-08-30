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
	cnt := 1
	for i := 0; i < n; i++ {
		if a[i] != cnt {
			ans++
			continue
		}
		cnt++
	}

	if ans <= n-1 {
		fmt.Fprintf(writer, "%v\n", ans)
	} else {
		fmt.Fprintf(writer, "%v\n", -1)
	}
}
