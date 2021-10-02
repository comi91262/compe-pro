package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func (q *queue) front() int {
	return (*q)[0]
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	a := queue{}
	b := queue{}
	c := queue{}

	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		a.push(tmp)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		b.push(tmp)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		c.push(tmp)
	}
	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)

	ans := 0
	for !a.empty() {
		s := a.pop()

		for !b.empty() && s >= b.front() {
			b.pop()
		}
		if b.empty() {
			break
		}
		t := b.pop()

		for !c.empty() && t >= c.front() {
			c.pop()
		}
		if c.empty() {
			break
		}
		c.pop()
		ans++
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
