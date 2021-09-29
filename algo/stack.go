package algo

type stack []int

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func (s *stack) pop() int {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}

func (q *stack) empty() bool {
	return len(*q) == 0
}
