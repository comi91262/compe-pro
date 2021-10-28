package datastructures

type Deque struct {
	cl   int
	cr   int
	data []int
}

func (d *Deque) PushFront(x int) {
	d.cl--
	d.data[d.cl] = x
}

func (d *Deque) PushBack(x int) {
	d.data[d.cr] = x
	d.cr++
}

func (d *Deque) PopFront() int {
	v := d.data[d.cl]
	d.cl++
	return v
}

func (d *Deque) PopBack() int {
	v := d.data[d.cr-1]
	d.cr--
	return v
}

func (d *Deque) Size() int {
	return d.cr - d.cl
}

func (d *Deque) Empty() bool {
	return d.Size() == 0
}

func (d *Deque) Get(x int) int {
	return d.data[d.cl+x-1]
}

func (d *Deque) Debug() (a []int) {
	a = make([]int, d.Size())
	for i := d.cl; i < d.cr; i++ {
		a[i-d.cl] = d.data[i]
	}
	return
}

func NewDeque(n int) *Deque {
	return &Deque{
		cl:   n,
		cr:   n,
		data: make([]int, n*2, n*2),
	}
}
