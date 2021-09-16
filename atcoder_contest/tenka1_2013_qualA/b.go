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

	cnt := 0
	for i := 0; i < n; i++ {
		var v int
		fmt.Fscan(reader, &v)
		var w int
		fmt.Fscan(reader, &w)
		var x int
		fmt.Fscan(reader, &x)
		var y int
		fmt.Fscan(reader, &y)
		var z int
		fmt.Fscan(reader, &z)
		s := v + w + x + y + z
		if 0 <= s && s < 20 {
			cnt++
		}
	}
	fmt.Fprintf(writer, "%v\n", cnt)
}
