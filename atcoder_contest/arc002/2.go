package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func isLeapYear(y int) bool {
	isOk := false
	isOk = y%4 == 0
	isOk = isOk && y%100 != 0
	isOk = isOk || y%400 == 0
	return isOk

	//t, s := false, false

	//if y%4 == 0 {
	//	t = true
	//	s = true
	//}

	//if y%100 == 0 {
	//	t = false
	//	s = true
	//}

	//if y%400 == 0 {
	//	t = true
	//	s = true
	//}

	//if !s {
	//	t = false
	//}

	//return t
}

func main() {
	defer writer.Flush()

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "/")
	a := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		a[i], _ = strconv.Atoi(s[i])
	}

	i := a[0]
	j := a[1]
	k := a[2]
	for ; ; i++ {
		for ; ; j++ {
			for ; ; k++ {
				if i%(j*k) == 0 {
					fmt.Fprintf(writer, "%v/%02v/%02v\n", i, j, k)
					return
				}
				if j == 2 && k == 28 && !isLeapYear(i) {
					break
				}
				if j == 2 && k == 29 && isLeapYear(i) {
					break
				}
				if (j == 4 || j == 6 || j == 9 || j == 11) && k == 30 {
					break
				}
				if k == 31 {
					break
				}
			}
			k = 1

			if j == 12 {
				break
			}
		}
		j = 1
	}
}
