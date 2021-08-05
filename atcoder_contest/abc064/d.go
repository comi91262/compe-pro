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

	ans := []string{}
	stack := []string{}
	m := map[string]int{}
	for i := range s {
		m[s[i]]++
		stack = append(stack, s[i])
		if m["("] < m[")"] {
			ans = append(ans, "(")
			m["("]++
		}
	}

	ans = append(ans, stack...)
	for i := 0; i < m["("]-m[")"]; i++ {
		ans = append(ans, ")")
	}

	fmt.Fprintf(writer, "%v\n", strings.Join(ans, ""))
}
