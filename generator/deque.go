package main

import (
	"fmt"
)

func main() {
	var T string
	fmt.Scanf("%v", &T)

	fmt.Printf(`
type %v struct {
}

type Deque struct {
	cl   int
	cr   int
	data []%v
}

func (d *Deque) pushFront(x %v) {
	d.cl--
	d.data[d.cl] = x
}

func (d *Deque) pushBack(x %v) {
	d.data[d.cr] = x
	d.cr++
}

func (d *Deque) popFront() %v {
	v := d.data[d.cl]
	d.cl++
	return v
}

func (d *Deque) popBack() %v {
	v := d.data[d.cr - 1]
	d.cr--
	return v
}

func (d *Deque) size() int {
	return d.cr - d.cl
}

func (d *Deque) empty() bool {
	return d.size() == 0
}

func (d *Deque) get(x int) %v {
	return d.data[d.cl + x - 1]
}

func create(size int) Deque {
	return Deque{cl: size, cr: size, data: make([]%v, size*2, size*2)}
}`, T, T, T, T, T, T, T, T)
}
