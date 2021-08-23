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

	var r int
	fmt.Fscan(reader, &r)
	var d int
	fmt.Fscan(reader, &d)
	var x int
	fmt.Fscan(reader, &x)

	ans := 0
	for i := 0; i < 10; i++ {
		if i == 0 {
			ans = r*x - d
		} else {
			ans = r*ans - d
		}
		fmt.Fprintf(writer, "%v\n", ans)
	}

}
