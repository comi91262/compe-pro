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
	var c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		d := []int{}
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if c[i]%c[j] == 0 {
				d = append(d, c[j])
			}
		}

		s := len(d)
		if s%2 == 0 {
			ans += float64(s+2) / float64(2*s+2)
		} else {
			ans += 1.0 / 2.0
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
