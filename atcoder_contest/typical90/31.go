package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func mex(a []int) int {
	sort.Ints(a)
	mx := 0
	for _, x := range a {
		if x < mx {
			continue
		}
		if x != mx {
			break
		}
		mx++
	}
	return mx
}

func grundy(w, b int) [][]int {
	g := make([][]int, w+1)
	for i := 0; i < w+1; i++ {
		g[i] = make([]int, b+w*(w+1)/2+1)
	}

	for i := 2; i < b+w*(w+1)/2+1; i++ {
		set := []int{}
		for j := 1; j <= i/2; j++ {
			set = append(set, g[0][i-j])
		}
		g[0][i] = mex(set)
	}

	// g[1][1] = mex([]int{g[0][2]})

	// (1, b) -> (0, b)
	// (w, b) -> (w-1, b+w) -> (0, b+w)
	for i := 1; i <= w; i++ {
		for j := 0; j < b+w*(w+1)/2+1-i; j++ {
			set := []int{}
			set = append(set, g[i-1][j+i])
			for k := 1; k <= j/2; k++ {
				set = append(set, g[i][j-k])
			}
			g[i][j] = mex(set)
		}
	}

	//	for i := 1; i < b+w*(w+1)/2+1; i++ {
	//		for j := 0; j < i && j <= w-1; j++ {
	//			g[j+1][i-j-1] = mex([]int{g[j][i-j]})
	//		}
	//	}

	for i := 0; i < w+1; i++ {
		// fmt.Fprintf(writer, "%v\n", g[i])
	}

	return g
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}

	g := grundy(50, 50)

	// var grundy [55][1555]int
	// for i := 0; i <= 50; i++ {
	// 	for j := 0; j <= 1500; j++ {
	// 		var mex [1555]int

	// 		if i >= 1 {
	// 			mex[grundy[i-1][j+i]] = 1
	// 		}
	// 		if j >= 2 {
	// 			for k := 1; k <= (j / 2); k++ {
	// 				mex[grundy[i][j-k]] = 1
	// 			}
	// 		}
	// 		for k := 0; k < 1555; k++ {
	// 			if mex[k] == 0 {
	// 				grundy[i][j] = k
	// 				break
	// 			}
	// 		}
	// 	}
	// }

	xor := 0
	for i := 0; i < n; i++ {
		// fmt.Fprintf(writer, "%v %v %v\n", a[i], b[i], g[a[i]][b[i]])
		xor ^= g[a[i]][b[i]]
	}
	// fmt.Fprintf(writer, "%v\n", xor)

	if xor == 0 {
		fmt.Fprintf(writer, "%v\n", "Second")
	} else {
		fmt.Fprintf(writer, "%v\n", "First")
	}

}
