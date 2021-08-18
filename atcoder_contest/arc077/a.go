package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type Deque struct {
	cl   int
	cr   int
	data []int
}

func (d *Deque) pushFront(x int) {
	d.cl--
	d.data[d.cl] = x
}

func (d *Deque) pushBack(x int) {
	d.data[d.cr] = x
	d.cr++
}

func (d *Deque) popFront() int {
	v := d.data[d.cl]
	d.cl++
	return v
}

func (d *Deque) popBack() int {
	v := d.data[d.cr-1]
	d.cr--
	return v
}

func (d *Deque) size() int {
	return d.cr - d.cl
}

func (d *Deque) empty() bool {
	return d.size() == 0
}

func (d *Deque) get(x int) int {
	return d.data[d.cl+x-1]
}

func create(size int) Deque {
	return Deque{cl: size, cr: size, data: make([]int, size*2, size*2)}
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := create(n)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			ans.pushFront(a[i])
		} else {
			ans.pushBack(a[i])
		}
	}

	if n%2 == 0 {
		for i := 0; i < n; i++ {
			fmt.Fprintf(writer, "%v", ans.popBack())
			if i != n-1 {
				fmt.Fprintf(writer, " ")
			}
		}
	} else {
		for i := 0; i < n; i++ {
			fmt.Fprintf(writer, "%v", ans.popFront())
			if i != n-1 {
				fmt.Fprintf(writer, " ")
			}
		}
	}
	fmt.Fprintf(writer, "\n")
}
