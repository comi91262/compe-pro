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

	var s int
	fmt.Fscan(reader, &s)
	var t int
	fmt.Fscan(reader, &t)

	ans := 0
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			for c := 0; c <= 100; c++ {
				if a+b+c <= s && a*b*c <= t {
					ans++
				}
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}

