package algo

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

type point struct {
	x int
	y int
}

type queue2 []point

func (q *queue2) push(n point) {
	*q = append(*q, n)
}

func (q *queue2) pop() point {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue2) empty() bool {
	return len(*q) == 0
}
