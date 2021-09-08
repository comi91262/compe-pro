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
	var l int
	fmt.Fscan(reader, &l)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	ans := 0
	cnt := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "+":
			cnt++
			if cnt >= l {
				ans++
				cnt = 0
			}
		case "-":
			cnt--
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
