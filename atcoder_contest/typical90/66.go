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

	var l = make([]int, n)
	var r = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &l[i], &r[i])
	}

	ans := 0.0
	for i := 0; i < n-1; i++ {

		for j := i + 1; j < n; j++ {
			sum := 0
			all := 0
			l1, r1 := l[i], r[i]
			l2, r2 := l[j], r[j]

			for o := l1; o <= r1; o++ {
				for p := l2; p <= r2; p++ {
					if o > p {

						sum++
					}
					all++
				}
			}

			ans += float64(sum) / float64(all)
		}

	}

	fmt.Fprintf(writer, "%v\n", ans)
}
