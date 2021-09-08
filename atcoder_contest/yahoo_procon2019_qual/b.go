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

	m := map[int]int{}
	for i := 0; i < 3; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		m[a]++
		m[b]++
	}
	for _, v := range m {
		if v == 3 {
			fmt.Fprintf(writer, "%v\n", "NO")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "YES")

}
