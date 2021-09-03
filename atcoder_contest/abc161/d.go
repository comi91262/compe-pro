package main

import (
	"bufio"
	"fmt"
	"os"
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
func main() {
	defer writer.Flush()
	var k int
	fmt.Fscan(reader, &k)

	q := queue{}
	for i := 1; i <= 9; i++ {
		q.push(i)
		k--
		if k <= 0 {
			fmt.Fprintf(writer, "%v\n", i)
			return
		}
	}

	ans := 0
	for !q.empty() {
		n := q.pop()

		a := 10*n + n%10 - 1
		b := 10*n + n%10
		c := 10*n + n%10 + 1

		if a%10 != 9 {
			q.push(a)
			k--
			if k <= 0 {
				ans = a
				break
			}
		}
		q.push(b)
		k--
		if k <= 0 {
			ans = b
			break
		}
		if c%10 != 0 {
			q.push(c)
			k--
			if k <= 0 {
				ans = c
				break
			}
		}

	}
	fmt.Fprintf(writer, "%v\n", ans)
}
