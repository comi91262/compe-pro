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
	for i := 0; i < n; i++ {
		for a[i] > 0 {
			if a[i]%2 == 0 || a[i]%3 == 2 {
				ans++
				a[i]--
			} else {
				break
			}
		}
		//		fmt.Fprintf(writer, "%v\n", ans)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
