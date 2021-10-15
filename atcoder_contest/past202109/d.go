package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func divisor(n int) []int {
	var r []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			r = append(r, i)
			if i*i != n {
				r = append(r, n/i)
			}
		}
	}
	return r
}

func main() {
	defer writer.Flush()
	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)

	xn := len(divisor(x))
	yn := len(divisor(y))
	switch {
	case xn > yn:
		fmt.Fprintf(writer, "%v\n", "X")
	case xn == yn:
		fmt.Fprintf(writer, "%v\n", "Z")
	case xn < yn:
		fmt.Fprintf(writer, "%v\n", "Y")
	}

}
