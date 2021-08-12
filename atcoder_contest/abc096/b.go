package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)
	var k int
	fmt.Fscan(reader, &k)

	d := []int{a, b, c}
	sort.Slice(d, func(i, j int) bool { return d[i] > d[j] })

	fmt.Fprintf(writer, "%v\n", d[0]*pow(2, k)+d[1]+d[2])

}
