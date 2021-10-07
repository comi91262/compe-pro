package main

import (
	"bufio"
	"fmt"
	"os"
)

var a [2000][2000]int
var b [2000][2000]int
var row, col [2000]int

func main() {
	r := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var h, w int
	fmt.Fscan(r, &h, &w)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fscan(r, &a[i][j])
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			row[i] += a[i][j]
			col[j] += a[i][j]
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			b[i][j] = row[i] + col[j] - a[i][j]
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fprint(writer, b[i][j])
			if j != w-1 {
				fmt.Fprint(writer, " ")
			}

		}
		if i != h-1 {
			fmt.Fprintln(writer, "")
		}
	}

}
package main

import (
	"bufio"
	. "bytes"
	"os"
	"strconv"
	"strings"
)

var wr = bufio.NewWriterSize(os.Stdout, 100000)

var l int
var bytes []byte

func scanInt() (result int) {
	for bytes[l] < '0' || bytes[l] > '9' {
		l++
	}
	for '0' <= bytes[l] && bytes[l] <= '9' {
		result = result*10 + int(bytes[l]-'0')
		l++
	}
	return result
}

func main() {
	defer wr.Flush()

	var a [2000][2000]int
	var row [2000]int
	var col [2000]int

	var buf Buffer
	buf.Grow(2000 * 2000)
	buf.ReadFrom(os.Stdin)
	bytes = buf.Bytes()

	h, w := scanInt(), scanInt()

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			a[i][j] = scanInt()
			row[i] += a[i][j]
			col[j] += a[i][j]
		}
	}

	for i := 0; i < h; i++ {
		if i > 0 {
			wr.WriteByte('\n')
		}
		ans := make([]string, w)
		for j := 0; j < w; j++ {
			ans[j] = strconv.Itoa(row[i] + col[j] - a[i][j])
		}
		wr.WriteString(strings.Join(ans, " "))
	}
}
