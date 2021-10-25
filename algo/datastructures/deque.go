package datastructures

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
