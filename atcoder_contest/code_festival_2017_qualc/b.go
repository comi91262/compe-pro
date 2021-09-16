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

	power := make([]int, n+1)
	base := 3
	power[0] = 1
	for i := 0; i < n; i++ {
		power[i+1] = power[i] * base
	}

	cnt := 0
	for i := 0; i < power[n]; i++ {
		prod := 1
		for j := 0; j < n; j++ {
			bit := i / power[j] % base
			switch bit {
			case 0:
				prod *= a[j]
			case 1:
				prod *= (a[j] + 1)
			case 2:
				prod *= (a[j] - 1)
			}
		}
		if prod%2 == 0 {
			cnt++
		}
	}
	fmt.Fprintf(writer, "%v\n", cnt)
}
