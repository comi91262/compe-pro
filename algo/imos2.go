package algo

func imos() {
	var a [][]int
	var n, w int

	for i := 0; i < n; i++ {
		var lx, ly, rx, ry int
		//		fmt.Fscan(reader, &lx, &ly, &rx, &ry)

		a[lx][ly]++
		a[lx][ry]--
		a[rx][ly]--
		a[rx][ry]++
	}

	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			a[i][j+1] += a[i][j]
		}
	}

	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			a[i+1][j] += a[i][j]
		}
	}
}
