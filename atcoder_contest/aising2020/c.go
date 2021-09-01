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

	m := map[int]int{}
	for x := 1; x <= 100; x++ {
		for y := 1; y <= 100; y++ {
			for z := 1; z <= 100; z++ {
				m[x*x+y*y+z*z+x*y+y*z+z*x]++
			}
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintf(writer, "%v\n", m[i])
	}
}
