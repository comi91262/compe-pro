package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func combination(n, k int) int {
	r := 1
	for d := 1; d <= k; d++ {
		r *= n
		n--
		r /= d

	}

	return r
}

func main() {
	defer writer.Flush()

	var l int
	fmt.Fscan(reader, &l)

	fmt.Fprintf(writer, "%v\n", combination(l-1, 11))

}
