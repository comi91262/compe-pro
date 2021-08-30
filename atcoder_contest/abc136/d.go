package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type queue []int

func (q *queue) push(n int) {
	*q = append(*q, n)
}

func (q *queue) pop() int {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) empty() bool {
	return len(*q) == 0
}
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
func main() {
	defer writer.Flush()
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	r := make([][]int, len(s))
	q := queue{}
	for i := 0; i < len(s); i++ {
		if s[i] == "R" {
			q.push(i)
		} else {
			for !q.empty() {
				r[i] = append(r[i], q.pop())
			}
		}
	}

	l := make([][]int, len(s))
	q = queue{}
	for i := 0; i < len(s); i++ {
		j := len(s) - 1 - i
		if s[j] == "L" {
			q.push(j)
		} else {
			for !q.empty() {
				l[j] = append(l[j], q.pop())
			}
		}
	}
	ans := make([]int, len(s))
	for i := 0; i < len(r); i++ {
		if len(r[i]) == 0 {
			continue
		}
		for j := range r[i] {
			if abs(i-r[i][j])%2 == 0 {
				ans[i]++
			} else {
				ans[i-1]++
			}
		}
	}
	for i := 0; i < len(l); i++ {
		if len(l[i]) == 0 {
			continue
		}
		for j := range l[i] {
			if abs(i-l[i][j])%2 == 0 {
				ans[i]++
			} else {
				ans[i+1]++
			}
		}
	}

	for i := range ans {
		fmt.Fprintf(writer, "%v", ans[i])
		if i != len(ans)-1 {
			fmt.Fprintf(writer, " ")
		}
	}
	fmt.Fprintf(writer, "\n")
}
