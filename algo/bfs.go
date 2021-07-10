package algo

func bfs2(h, w int, sx, sy, ex, ey int, s [][]string) {
	const inf = 1 << 60
	var dx = [4]int{1, 0, -1, 0}
	var dy = [4]int{0, 1, 0, -1}

	var dist = make([][]int, h)
	for i := 0; i < h; i++ {
		dist[i] = make([]int, w)
		for j := 0; j < w; j++ {
			dist[i][j] = inf
			dist[i][j] = inf
			dist[i][j] = inf
			dist[i][j] = inf
		}
	}

	deq := create(4080000)
	dist[sx][sy] = 0
	deq.pushFront(state{sx, sy})

	for !deq.empty() {
		u := deq.popFront()
		for i := 0; i < 4; i++ {
			tx := u.x + dx[i]
			ty := u.y + dy[i]
			cost := dist[u.x][u.y]
			if 0 <= tx && tx < h && 0 <= ty && ty < w && s[tx][ty] == "." && dist[tx][ty] > cost {
				dist[tx][ty] = cost
				// if u != i {
				// 	deq.pushBack(state{tx, ty, i})
				// } else {
				// 	deq.pushFront(state{tx, ty, i})
				// }
			}
		}
	}
}

type state struct {
	x int
	y int
}

type Deque struct {
	cl   int
	cr   int
	data []state
}

func (d *Deque) pushFront(x state) {
	d.cl--
	d.data[d.cl] = x
}

func (d *Deque) pushBack(x state) {
	d.data[d.cr] = x
	d.cr++
}

func (d *Deque) popFront() state {
	v := d.data[d.cl]
	d.cl++
	return v
}

func (d *Deque) popBack() state {
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

func (d *Deque) get(x int) state {
	return d.data[d.cl+x-1]
}

func create(size int) Deque {
	return Deque{cl: size, cr: size, data: make([]state, size*2, size*2)}
}
