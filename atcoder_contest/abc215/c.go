package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func nextPermutation(a []string) {
	var lastIndex = len(a) - 1

	if lastIndex < 1 {
		return
	}

	var i = lastIndex - 1
	for i >= 0 && a[i] >= a[i+1] {
		i -= 1
	}
	if i < 0 {
		reverse(a, 0, lastIndex)
		return
	}

	var j = lastIndex
	for j > i+1 && a[j] <= a[i] {
		j -= 1
	}

	swap(a, i, j)
	reverse(a, i+1, lastIndex)
}

func main() {
	defer writer.Flush()

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	sort.Strings(s)

	var k int
	fmt.Fscan(reader, &k)

	for i := 0; i < k-1; i++ {
		nextPermutation(s)
	}
	for i := 0; i < len(s); i++ {
		fmt.Fprintf(writer, "%v", s[i])
	}
	fmt.Fprintf(writer, "\n")

}
