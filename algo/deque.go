package algo

type Deque struct {
	cl   int
	cr   int
	data []int
}

func (d *Deque) addFront(x int) {
	d.cl--
	d.data[d.cl] = x
}

func (d *Deque) addBack(x int) {
	d.data[d.cr] = x
	d.cr++
}

func (d *Deque) removeFront() {
	d.cl++
}

func (d *Deque) removeBack() {
	d.cr--
}

func (d *Deque) get(x int) int {
	return d.data[d.cl+x-1]
}

func create(size int) Deque {
	return Deque{cl: size / 2, cr: size / 2, data: make([]int, size)}
}
