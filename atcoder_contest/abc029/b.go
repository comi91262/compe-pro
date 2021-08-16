package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	n := 12
	var a = make([][]string, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		a[i] = strings.Split(tmp, "")
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] == "r" {
				ans++
				break
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
