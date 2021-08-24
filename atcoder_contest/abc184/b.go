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
	var n int
	fmt.Fscan(reader, &n)
	var x int
	fmt.Fscan(reader, &x)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	ans := x
	for i := 0; i < n; i++ {
		switch s[i] {
		case "o":
			ans++
		case "x":
			if ans > 0 {
				ans--
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
