package algo

type PrefixSum2 struct {
	row int
	col int
	a   [][]int
}

func NewPrefixSum(n, m int) *PrefixSum2 {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
	}

	return &PrefixSum2{n, m, a}
}

func (p *PrefixSum2) Add(x, y, v int) {
	p.a[x][y] += v
}

func (p *PrefixSum2) Build() {
	for i := 0; i < p.row; i++ {
		for j := 1; j < p.col; j++ {
			p.a[i][j] += p.a[i][j-1]
		}
	}

	for i := 0; i < p.col; i++ {
		for j := 1; j < p.row; j++ {
			p.a[j][i] += p.a[j-1][i]
		}
	}
}

func (p *PrefixSum2) Get(sx, sy, tx, ty int) (result int) {
	if tx < 0 || ty < 0 {
		return
	}
	result = p.a[tx][ty]
	if sy >= 0 {
		result -= p.a[tx][sy]
	}
	if sx >= 0 {
		result -= p.a[sx][ty]
	}
	if sx >= 0 && sy >= 0 {
		result += p.a[sx][sy]
	}
	return
}
