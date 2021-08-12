package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type queue []string

func (q *queue) push(n string) {
	*q = append(*q, n)
}

func (q *queue) pop() string {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) empty() bool {
	return len(*q) == 0
}
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}
func main() {
	defer writer.Flush()
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	q := queue{}

	ans := 0
	for i := range s {
		q.push(s[i])

		for !q.empty() && (s[i] != "A" && s[i] != "C" && s[i] != "G" && s[i] != "T") {
			q.pop()
		}

		ans = max(ans, len(q))
	}
	fmt.Fprintf(writer, "%v\n", ans)
}

