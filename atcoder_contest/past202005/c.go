package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func pow(ini, a, x int) (int, bool) {
	r := ini
	for x > 0 {
		if x&1 == 1 {
			if r > 1_000_000_000/a {
				return 0, false
			}
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r, true
}
func main() {
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)
	var r int
	fmt.Fscan(reader, &r)
	var n int
	fmt.Fscan(reader, &n)

	ans, isOk := pow(a, r, n-1)
	if isOk {
		fmt.Fprintf(writer, "%v\n", ans)
	} else {
		fmt.Fprintf(writer, "%v\n", "large")
	}
}
