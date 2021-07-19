package algo

// 全点対間最短経路を求めるワーシャルフロイド法 O(n * n * n)
// d[i][j]: i -> j 間の距離
func warshallFloyd(d [][]int, n int) [][]int {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	return d
}
