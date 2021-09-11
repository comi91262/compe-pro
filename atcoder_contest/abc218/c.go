package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

// in-place  n * n
func rotateMatrix(a [][]string) {
	n := len(a)

	for x := 0; x < n/2; x++ {
		for y := x; y < n-x-1; y++ {
			a[x][y], a[y][n-1-x], a[n-1-x][n-1-y], a[n-1-y][x] = a[y][n-1-x], a[n-1-x][n-1-y], a[n-1-y][x], a[x][y]
		}
	}
}

func diffMatrix(s [][]string, t [][]string) (int, int) {
	n := len(s)

	sx, sy := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == "#" {
				sx, sy = i, j
			}
		}

	}
	tx, ty := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if t[i][j] == "#" {
				tx, ty = i, j
			}
		}
	}

	return tx - sx, ty - sy
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	s := make([][]string, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	t := make([][]string, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		t[i] = strings.Split(tmp, "")
	}

	sn := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == "#" {
				sn++
			}
		}
	}
	tn := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if t[i][j] == "#" {
				tn++
			}
		}
	}
	if tn != sn {
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}

	for i := 0; i < 4; i++ {
		dx, dy := diffMatrix(s, t)
		isOk := true
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if s[j][k] == "#" {
					nx := j + dx
					ny := k + dy
					if nx < 0 || nx >= n {
						isOk = false
						break
					}
					if ny < 0 || ny >= n {
						isOk = false
						break
					}

					if t[nx][ny] != "#" {
						isOk = false
						break
					}
				}

			}

		}
		if isOk {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}

		rotateMatrix(s)
	}

	fmt.Fprintf(writer, "%v\n", "No")

}
