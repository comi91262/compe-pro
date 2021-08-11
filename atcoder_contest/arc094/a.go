package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)

	d := []int{a, b, c}
	sort.Ints(d)

	b1 := d[2] - d[0]
	b2 := d[2] - d[1]

	ans := 0
	if b1%2 == 0 {
		ans += b1 / 2
		if b2%2 == 0 {
			ans += b2 / 2
		} else {
			ans += b2 / 2
			ans += 2
		}
	} else {
		ans += b1 / 2
		if b2%2 == 0 {
			ans += b2 / 2
			ans += 2
		} else {
			ans += b2 / 2
			ans += 1
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
