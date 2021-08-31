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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var ss string
	fmt.Fscan(reader, &ss)
	var c = strings.Split(ss, "")

	m := map[string]int{}
	for i := 0; i < n; i++ {
		m[c[i]]++
	}
	rn := m["R"]

	q := queue{}

	for i := 0; i < n; i++ {
		if i >= rn {
			q.push(c[i])
		}
	}
	ans := 0
	for i := 0; i < rn; i++ {
		if c[i] == "R" {
			continue
		}
		if !q.empty() {
			q.pop()
			ans++
		} else {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
