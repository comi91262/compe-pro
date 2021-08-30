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

	h := 0
	for i := n; i > 0; i /= 2 {
		h++
	}

	cnt := 0
	m := 1
	if h%2 == 1 {
		for m <= n {
			if cnt%2 == 0 {
				m = 2*m + 1
			} else {
				m = 2 * m
			}
			cnt++
		}
	} else {
		for m <= n {
			if cnt%2 == 0 {
				m = 2 * m
			} else {
				m = 2*m + 1
			}
			cnt++
		}
	}

	if cnt%2 == 1 {
		fmt.Fprintf(writer, "%v\n", "Aoki")
	} else {
		fmt.Fprintf(writer, "%v\n", "Takahashi")
	}
}
