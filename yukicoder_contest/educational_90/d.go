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

	base := 3

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	power := make([]int, n+1)
	power[0] = 1
	for i := 0; i < n; i++ {
		power[i+1] = power[i] * base
	}

	for i := 0; i < power[n]; i++ {

		b, c, d := 0, 0, 0
		for j := 0; j < n; j++ {
			bit := i / power[j] % base
			switch bit {
			case 0:
				b += a[j]
			case 1:
				c += a[j]
			case 2:
				d += a[j]
			}
		}

		if b == c && c == d {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "No")
}
