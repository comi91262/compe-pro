package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func toNumber(a []int, base int) int {
	cnt := 1
	ans := 0
	for i := 0; i < len(a); i++ {
		ans += a[len(a)-1-i] * cnt
		cnt *= base
	}
	return ans
}

func main() {
	defer writer.Flush()

	var k int
	fmt.Fscan(reader, &k)
	var ss string
	fmt.Fscan(reader, &ss)
	var a = strings.Split(ss, "")
	fmt.Fscan(reader, &ss)
	var b = strings.Split(ss, "")

	c := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = int(a[i][0] - "0"[0])
	}
	d := make([]int, len(b))
	for i := 0; i < len(b); i++ {
		d[i] = int(b[i][0] - "0"[0])
	}

	e := toNumber(c, k)
	f := toNumber(d, k)

	fmt.Fprintf(writer, "%v\n", e*f)
}
