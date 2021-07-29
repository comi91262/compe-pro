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
	var k int
	fmt.Fscan(reader, &k)
	var s int
	fmt.Fscan(reader, &s)

	cnt := 0
	for i := 0; i <= k; i++ {
		for j := 0; j <= k; j++ {
			z := s - i - j
			if 0 <= z && z <= k {
				cnt++
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", cnt)
}
