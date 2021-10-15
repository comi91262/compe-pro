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
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	m := []int{}
	for i := 0; i < n; i++ {
		if s[i] == "0" {
			m = append(m, i)
		}
	}

	if len(m) == 1 {
		fmt.Fprintf(writer, "%v\n", -1)
		return
	}

	ans := make([]int, n)
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == "1" {
			ans[i] = i + 1
			continue
		}

		ans[i] = m[(cnt+1)%len(m)] + 1
		cnt++
	}

	for i := range ans {
		if i > 0 {
			fmt.Fprintf(writer, " ")
		}
		fmt.Fprintf(writer, "%v", ans[i])
	}
	fmt.Fprintf(writer, "\n")
}
