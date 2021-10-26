package algo

func chmin(x *int, y int) {
	*x = min(*x, y)
}
func chmax(x *int, y int) {
	*x = max(*x, y)
}

// 最小, 配る, 一次元DP
// 階段を登るDP
func flog1(n int, h []int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 << 60
	}

	dp[0] = 0
	for i := 0; i < n; i++ {
		if i+1 < n {
			chmin(&dp[i+1], dp[i]+abs(h[i]-h[i+1]))
		}
		if i+2 < n {
			chmin(&dp[i+2], dp[i]+abs(h[i]-h[i+2]))
		}
	}
	return dp[n-1]
}

// 最小, 配る, 一次元DP
// 階段を登るDP
func flog2(n, k int, h []int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 << 60
	}

	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := 1; j <= k; j++ {
			if i+j < n {
				chmin(&dp[i+j], dp[i]+abs(h[i]-h[i+j]))
			}
		}
	}
	return dp[n-1]
}

// 条件に従って、変数を追加する系DP
func vacation(n int, a, b, c []int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 3)
	}

	for i := 0; i < n; i++ {
		dp[i+1][0] += max(dp[i][1]+a[i], dp[i][2]+a[i])
		dp[i+1][1] += max(dp[i][0]+b[i], dp[i][2]+b[i])
		dp[i+1][2] += max(dp[i][0]+c[i], dp[i][1]+c[i])
	}
	return max(dp[n]...)
}

// ナップサックDP
func knapsack1(n, mw int, w, v []int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, mw+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= mw; j++ {
			chmax(&dp[i+1][j], dp[i][j])
			if j+w[i] <= mw {
				chmax(&dp[i+1][j+w[i]], dp[i][j]+v[i])
			}
		}
	}
	return dp[n][mw]
}

// ナップサックDP 亜種 価値の最大化ではなく重さの最小化
func knapsack2(n, mw int, w, v []int) int {
	dp := make([][]int, n+1)
	mv := 100001
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, mv)
		for j := 0; j < mv; j++ {
			dp[i][j] = 1 << 60
		}
	}

	dp[0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j < mv; j++ {
			chmin(&dp[i+1][j], dp[i][j])
			if j+v[i] < mv {
				chmin(&dp[i+1][j+v[i]], dp[i][j]+w[i])
			}
		}
	}
	value := 0
	for i := 0; i < mv; i++ {
		if dp[n][i] <= mw {
			chmax(&value, i)
		}
	}
	return value
}

// LCS
func lcs(s, t string) string {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				chmax(&dp[i+1][j+1], dp[i][j]+1)
			} else {
				chmax(&dp[i+1][j+1], max(dp[i+1][j], dp[i][j+1]))
			}
		}
	}

	ans := make([]rune, dp[len(s)][len(t)])
	i, j, l := len(s), len(t), len(ans)
	for l > 0 {
		switch dp[i][j] {
		case dp[i-1][j]:
			i--
		case dp[i][j-1]:
			j--
		case dp[i-1][j-1] + 1:
			ans[l-1] = rune(s[i-1])
			l--
			i--
			j--
		}
	}
	return string(ans)
}

// 経路計算DP
func gred1(h, w int, s [][]string) int {
	dp := make([][]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, w)
	}

	dp[0][0] = 1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i+1 < h && s[i+1][j] != "#" {
				dp[i+1][j] += dp[i][j]
			}
			if j+1 < w && s[i][j+1] != "#" {
				dp[i][j+1] += dp[i][j]
			}
		}
	}
	return dp[h-1][w-1]
}

//  確率DP
func coins(n int, p []float64) float64 {
	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, n+1)
		for j := 0; j <= n; j++ {
		}
	}

	dp[0][0] = 1.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j] += dp[i][j] * (1.0 - p[i])
			dp[i+1][j+1] += dp[i][j] * p[i]
		}
	}

	pro := 0.0
	for i := 0; i <= n; i++ {
		if n-i < i {
			pro += dp[n][i]
		}
	}
	return pro
}

// 後退解析DP
func Stones(n, k int, a []int) bool {
	dp := make([]bool, k+1)

	dp[0] = false
	for i := 0; i <= k; i++ {
		for j := 0; j < n; j++ {
			if i-a[j] >= 0 && !dp[i-a[j]] {
				dp[i] = true
			}
		}
	}
	return dp[k]
}

// DFSによる木DP
func treeDp(n, pre int) {
	var dp [100001]int
	var g [][]int
	dp[n]++

	for _, next := range g[n] {
		if pre == next {
			continue
		}

		treeDp(next, n)
		dp[n] += dp[next]
	}

}

// 部分和DP
func subSumDp() {
	var dp [101][100001]bool
	var n, sum int
	var t []int

	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < sum+1; j++ {
			if j >= t[i] {
				dp[i+1][j] = dp[i][j-t[i]] || dp[i][j]
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
}

// 区間DP dp[l][r] = min(dp[l][k]+dp[k+1][r], dp[l][r]) or min(dp[l+1][r-1]+v, dp[l][r])
func intervalDp(n int, a []int) {
	inf := 1_000_000_000 + 7
	dp := make([][]int, 2*n)
	for i := 0; i < 2*n; i++ {
		dp[i] = make([]int, 2*n)
	}

	for i := 0; i < 2*n; i++ {
		for j := 0; j < 2*n; j++ {
			dp[i][j] = inf
		}
	}
	for i := 0; i < 2*n-1; i++ {
		dp[i][i+1] = abs(a[i] - a[i+1])
	}

	for interval := 3; interval <= 2*n; interval += 2 {
		for j := 0; j < 2*n-interval; j++ {
			cl, cr := j, j+interval

			dp[cl][cr] = min(dp[cl][cr], dp[cl+1][cr-1]+abs(a[cl]-a[cr]))

			for k := cl; k < cr; k++ {
				dp[cl][cr] = min(dp[cl][cr], dp[cl][k]+dp[k+1][cr])
			}
		}
	}
}

func bitDp(n, k int) {
	cost := make([]int, 1<<n)
	for i := 1; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			for l := 0; l < j; l++ {
				if ((i>>j)&1) == 1 && ((i>>l)&1) == 1 {
				}
			}
		}
	}

	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, 1<<n)
		for j := 0; j < 1<<n; j++ {
			dp[i][j] = 1 << 62
		}
	}
	dp[0][0] = 0
	for i := 1; i <= k; i++ {
		for bit := 1; bit < 1<<n; bit++ {
			for b := bit; b != 0; b = (b - 1) & bit {
				dp[i][bit] = min(dp[i][bit], max(dp[i-1][bit-b], cost[b]))
			}
		}
	}
}
