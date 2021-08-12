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

	a := []int{111, 222, 333, 444, 555, 666, 777, 888, 999}

	ans := 0
	mn := 1 << 60
	for i := range a {
		if mn >= a[i]-n && a[i]-n >= 0 {
			mn = a[i] - n
			ans = a[i]
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
