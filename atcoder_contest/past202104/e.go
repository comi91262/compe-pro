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
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	d := create(3 * 100000)

	for i := 0; i < n; i++ {
		switch s[i] {
		case "L":
			d.pushFront(i + 1)
		case "R":
			d.pushBack(i + 1)
		case "A":
			if d.size() <= 0 {
				fmt.Fprintf(writer, "%v\n", "ERROR")
				break
			}
			fmt.Fprintf(writer, "%v\n", d.popFront())
		case "B":
			if d.size() <= 1 {
				fmt.Fprintf(writer, "%v\n", "ERROR")
				break
			}
			v := d.popFront()
			fmt.Fprintf(writer, "%v\n", d.popFront())
			d.pushFront(v)
		case "C":
			if d.size() <= 2 {
				fmt.Fprintf(writer, "%v\n", "ERROR")
				break
			}
			v1 := d.popFront()
			v2 := d.popFront()
			fmt.Fprintf(writer, "%v\n", d.popFront())
			d.pushFront(v2)
			d.pushFront(v1)
		case "D":
			if d.size() <= 0 {
				fmt.Fprintf(writer, "%v\n", "ERROR")
				break
			}
			fmt.Fprintf(writer, "%v\n", d.popBack())
		case "E":
			if d.size() <= 1 {
				fmt.Fprintf(writer, "%v\n", "ERROR")
				break
			}
			v := d.popBack()
			fmt.Fprintf(writer, "%v\n", d.popBack())
			d.pushBack(v)
		case "F":
			if d.size() <= 2 {
				fmt.Fprintf(writer, "%v\n", "ERROR")
				break
			}
			v1 := d.popBack()
			v2 := d.popBack()
			fmt.Fprintf(writer, "%v\n", d.popBack())
			d.pushBack(v2)
			d.pushBack(v1)
		}
	}
}
