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
	var s string
	fmt.Fscan(reader, &s)

	ss := strings.Split(s, "+")

	ans := 0
	for i := range ss {
		sss := strings.Split(ss[i], "*")

		flag := true
		for j := range sss {
			if sss[j] == "0" {
				flag = false
			}
		}
		if flag {
			ans++
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
