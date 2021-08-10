package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func swap(a []string, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// スライス a の i番目からj番目を反転させる関数
func reverse(a []string, i, j int) {

	for k := 0; k < (j-i+1)/2; k++ {
		swap(a, i+k, j-k)
	}
}

func main() {
	defer writer.Flush()

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		reverse(s, l-1, r-1)
	}
	fmt.Fprintf(writer, "%v\n", strings.Join(s, ""))

}
