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
	var k int
	fmt.Fscan(reader, &k)

	t := make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &t[i][j])
		}
	}

	base := k
	power := make([]int, n+1)
	power[0] = 1
	for i := 0; i < n; i++ {
		power[i+1] = power[i] * base
	}

	for i := 0; i < power[n]; i++ {
		check := 0
		for j := 0; j < n; j++ {
			bit := i / power[j] % base
			check ^= t[j][bit]
		}

		if check == 0 {
			fmt.Fprintf(writer, "%v\n", "Found")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "Nothing")
}
