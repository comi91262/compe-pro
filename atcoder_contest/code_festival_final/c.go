package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
}

func main() {
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)

	for k := 10; k <= 10000; k++ {
		b := a
		c := ""
		for b > 0 {
			c = strconv.Itoa(b%k) + c
			b /= k
		}

		if n, _ := strconv.Atoi(c); k == n {
			fmt.Fprintf(writer, "%v\n", k)
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", -1)
}
