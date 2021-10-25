package graph

type edge struct {
	to   int
	cost int
}

type queue []int

func (q *queue) push(n int) {
	*q = append(*q, n)
}

func (q *queue) pop() int {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) front() int {
	return (*q)[0]
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func topoSort(g [][]edge) (sorted []int) {
	n := len(g)
	index := make([]int, n)
	for i := 0; i < n; i++ {
		for _, e := range g[i] {
			index[e.to]++
		}
	}

	q := queue{}
	for i := 0; i < n; i++ {
		if index[i] == 0 {
			q.push(i)
		}
	}

	for !q.empty() {
		now := q.pop()
		sorted = append(sorted, now)
		for _, e := range g[now] {
			index[e.to]--
			if index[e.to] == 0 {
				q.push(e.to)
			}
		}
	}
	return
}
