package graph

type node struct {
	num  int
	dist int
}

func (i node) Less(j QItem) bool {
	return i.dist < j.(node).dist
}

func dijkstra(start int, g [][]edge) []int {
	const inf = 1 << 60
	queue := NewPriorityQueue()

	dist := make([]int, len(g))
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[start] = 0
	queue.Push(node{start, dist[start]})

	for !queue.Empty() {
		v := queue.Pop().(node)
		if dist[v.num] < v.dist {
			continue
		}
		for _, e := range g[v.num] {
			if dist[e.to] > dist[v.num]+e.cost {
				dist[e.to] = dist[v.num] + e.cost
				queue.Push(node{e.to, dist[e.to]})
			}
		}
	}
	return dist
}

// nは頂点数、sは開始頂点
func bellmanFord(n, s int, g [][]edge) bool {
	inf := 1 << 60
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}

	dist[s] = 0 // 開始点の距離は0
	for i := 0; i < n; i++ {
		for v := 0; v < n; v++ {
			for _, e := range g[v] {
				if dist[v] != inf && dist[e.to] > dist[v]+e.cost {
					dist[e.to] = dist[v] + e.cost
					if i == n-1 {
						return true
					}
				}
			}
		}
	}
	return false
}

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
