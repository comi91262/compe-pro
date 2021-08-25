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
	var x int
	fmt.Fscan(reader, &x)

	for {
		flag := false
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				flag = true
				break
			}
		}
		if !flag {
			break
		}
		x++
	}
	fmt.Fprintf(writer, "%v\n", x)

}
