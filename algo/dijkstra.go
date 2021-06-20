package algo

import "math"

// fibonacci heap implemeted by https://github.com/theodesp/go-heaps
type Item interface {
	Compare(than Item) int
}

type FibonacciHeap struct {
	root *node
}
type node struct {
	item                      Item
	prev, next, parent, child *node
	isMarked                  bool
	degree                    int
}

func (fh *FibonacciHeap) Insert(item Item) Item {
	n := &node{item: item, isMarked: false}

	fh.insertRoot(n)
	return item
}

func (fh *FibonacciHeap) FindMin() Item {
	if fh.root == nil {
		return nil
	}
	return fh.root.item
}

func (fh *FibonacciHeap) DeleteMin() Item {
	r := fh.root
	if r == nil {
		return nil
	}
	for {
		// add r children to fh's root list
		if x := r.child; x != nil {
			x.parent = nil
			if x.next != x {
				r.child = x.next
				x.next.prev = x.prev
				x.prev.next = x.next
			} else {
				r.child = nil
			}
			x.prev = r.prev
			x.next = r
			r.prev.next = x
			r.prev = x
		} else {
			break
		}
	}
	// remove r from fh's root list
	r.prev.next = r.next
	r.next.prev = r.prev

	if r == r.next {
		fh.root = nil
	} else {
		fh.root = r.next
		fh.consolidate()
	}

	return r.item
}

func (fh *FibonacciHeap) consolidate() {
	degreeToRoot := make(map[int]*node)
	w := fh.root
	last := w.prev
	for {
		r := w.next
		x := w
		d := x.degree
		for {
			if y, ok := degreeToRoot[d]; !ok {
				break
			} else {
				if y.item.Compare(x.item) < 0 {
					y, x = x, y
				}
				link(x, y)
				delete(degreeToRoot, d)
				d++
			}
		}
		degreeToRoot[d] = x
		if w == last {
			break
		}
		w = r
	}
	fh.root = nil
	for _, v := range degreeToRoot {
		fh.insertRoot(v)
	}

}

func link(x, y *node) {
	// remove y from fh's root list
	y.next.prev = y.prev
	y.prev.next = y.next
	// make y a child of x and increase degree of x
	y.parent = x
	if x.child == nil {
		x.child = y
		y.prev = y
		y.next = y
	} else {
		insert(x.child, y)
	}

	y.isMarked = false
}

func (fh *FibonacciHeap) insertRoot(n *node) {
	if fh.root == nil {
		// create fh's root list containing only n
		n.prev = n
		n.next = n
		fh.root = n
	} else {
		// insert n to fh's root list
		insert(fh.root, n)
		if n.item.Compare(fh.root.item) < 0 {
			fh.root = n
		}
	}
}

func insert(x, y *node) {
	x.prev.next = y
	y.next = x
	y.prev = x.prev
	x.prev = y
}

type Edge struct {
	to   int
	cost int
}

type Vertex struct {
	value    int
	priority int
}

func (a Vertex) Compare(b Item) int {
	a1 := a
	a2 := b.(Vertex)
	switch {
	case a1.priority > a2.priority:
		return 1
	case a1.priority < a2.priority:
		return -1
	default:
		return 0
	}
}

func dijkstra(start int, g *[][]Edge, d *[]int) {
	for i := 0; i < len((*d)); i++ {
		(*d)[i] = math.MaxInt64
	}
	(*d)[start] = 0

	heap := &FibonacciHeap{}
	heap.Insert(Vertex{value: start, priority: 0})

	for heap.FindMin() != nil {
		v := heap.DeleteMin().(Vertex).value
		for _, e := range (*g)[v] {
			if (*d)[e.to] > (*d)[v]+e.cost {
				(*d)[e.to] = (*d)[v] + e.cost
				heap.Insert(Vertex{value: e.to, priority: (*d)[e.to]})
			}
		}
	}
}

//	for i := 0; i < m; i++ {
//		g[a] = append(g[a], Edge{to: b, cost: c})
//		g[b] = append(g[b], Edge{to: a, cost: c})
//	}
//
//	dijkstra(1)
