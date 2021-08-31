package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type Deque struct {
	cl   int
	cr   int
	data []string
}

func (d *Deque) pushFront(x string) {
	d.cl--
	d.data[d.cl] = x
}

func (d *Deque) pushBack(x string) {
	d.data[d.cr] = x
	d.cr++
}

func (d *Deque) popFront() string {
	v := d.data[d.cl]
	d.cl++
	return v
}

func (d *Deque) popBack() string {
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

func (d *Deque) get(x int) string {
	return d.data[d.cl+x-1]
}

func create(size int) Deque {
	return Deque{cl: size, cr: size, data: make([]string, size*2, size*2)}
}

func main() {
	defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)
	var q int
	fmt.Fscan(reader, &q)

	d := create(len(s) + q + 1)
	d.pushFront(s)
	p := true
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		switch t {
		case 1:
			p = !p
		case 2:
			var f int
			var c string
			fmt.Fscan(reader, &f, &c)
			switch f {
			case 1:
				if p {
					d.pushFront(c)
				} else {
					d.pushBack(c)
				}
			case 2:
				if p {
					d.pushBack(c)
				} else {
					d.pushFront(c)
				}
			}
		}
	}

	ans := strings.Join(d.data, "")
	if p {
		fmt.Fprintf(writer, "%v\n", ans)
	} else {
		a := []rune(ans)
		n := len(a)
		for i := 0; i < n/2; i++ {
			a[i], a[n-i-1] = a[n-i-1], a[i]
		}
		fmt.Fprintf(writer, "%v\n", string(a))
	}

}
