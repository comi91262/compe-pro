package algo

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// i文字目以降(i文字目を含む)に初めてj(0-9)が出現するのは何文字目かを管理するテーブル
// 0は見つからないことを示すため、1-origin
type NextAppearTable struct {
	dict [][]int
}

func NewNextAppperTable(s string) *NextAppearTable {
	table := NextAppearTable{}

	n := len(s)
	table.dict = make([][]int, n)
	for i := 0; i < n; i++ {
		table.dict[i] = make([]int, 10)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 10; j++ {
			if j == int(s[n-i-1]-'0') {
				table.dict[n-i-1][j] = n - i
				continue
			}

			if i > 0 {
				table.dict[n-i-1][j] = table.dict[n-i][j]
			}
		}
	}
	return &table
}

func (table *NextAppearTable) nextCur(cur, c int) int {
	return table.dict[cur][c]
}
