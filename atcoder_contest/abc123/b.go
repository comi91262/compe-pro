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
	var d int
	fmt.Fscan(reader, &d)
	var e int
	fmt.Fscan(reader, &e)

	f := []int{a, b, c, d, e}
	sort.Slice(f, func(i, j int) bool {
		return (f[i]-1)%10 > (f[j]-1)%10
	})

	ans := 0
	for i := 0; i < len(f)-1; i++ {
		ans += (f[i] + 9) / 10 * 10
	}

	fmt.Fprintf(writer, "%v\n", ans+f[len(f)-1])
}
