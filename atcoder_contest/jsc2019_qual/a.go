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
	var m int
	fmt.Fscan(reader, &m)
	var d int
	fmt.Fscan(reader, &d)

	day := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= d; j++ {
			d1 := j % 10
			d10 := j / 10 % 10
			if d1 >= 2 && d10 >= 2 && d1*d10 == i {
				day++
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", day)
}
