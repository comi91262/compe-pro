package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type food struct {
	hp     int
	attack int
	idx    int
}

type queue []food

func (q *queue) push(n food) {
	*q = append(*q, n)
}

func (q *queue) pop() food {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) front() food {
	return (*q)[0]
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var food = make([]food, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &food[i].hp, &food[i].attack)
		food[i].idx = i + 1
	}
	q := queue{}
	for i := 0; i < n; i++ {
		q.push(food[i])
	}
	if len(q) == 1 {
		fmt.Fprintf(writer, "%v\n", 1)
		return
	}
	for len(q) > 1 {
		u := q.pop()
		v := q.pop()

		s := (u.hp + v.attack - 1) / v.attack
		t := (v.hp + u.attack - 1) / u.attack

		if s == t && q.empty() {
			fmt.Fprintf(writer, "%v\n", -1)
			return
		}

		if s > t {
			q.push(u)
		} else if s == t {
		} else {
			q.push(v)
		}
	}
	winner := q.pop()
	for i := 0; i < n; i++ {
		if winner.idx == i+1 {
			continue
		}
		s := (winner.hp + food[i].attack - 1) / food[i].attack
		t := (food[i].hp + winner.attack - 1) / winner.attack

		if s <= t {
			fmt.Fprintf(writer, "%v\n", -1)
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", winner.idx)
}
