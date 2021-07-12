package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	var n int
	fmt.Fscan(reader, &n)

	var tmp string
	fmt.Fscan(reader, &tmp)
	s := strings.Split(tmp, "")

	d := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "a":
			d[i] = 0
		case "b":
			d[i] = 1
		case "c":
			d[i] = 2
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += pow(2, i-1) * d[i-1]
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
