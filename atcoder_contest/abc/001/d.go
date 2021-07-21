package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func byteToDigit(c byte) (int, error) {
	n := c - 48

	if n < 0 || 9 < n {
		return int(c), errors.New("wrong number")
	}

	return int(n), nil
}

func strToDigits(s string) []int {
	r := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		r[i], _ = byteToDigit(s[i])
	}
	return r
}

func toSecondFirst(time []int) int {
	if 0 <= time[3] && time[3] < 5 {
		time[3] = 0
	} else if 5 <= time[3] && time[3] < 10 {
		time[3] = 5
	}

	return (time[0]*10+time[1])*60 + time[2]*10 + time[3]
}

func toSecondSecond(time []int) int {
	switch {
	case 0 < time[3] && time[3] <= 5: // ok
		time[3] = 5
	case 5 < time[3] && time[3] <= 9:
		time[3] = 0

		switch {
		case time[2] != 5: // ok
			time[2]++
		case time[2] == 5:
			time[2] = 0

			switch {
			case time[1] != 9: // ok
				time[1]++
			case time[1] == 9:
				time[1] = 0
				time[0]++
			}
		}
	}

	return (time[0]*10+time[1])*60 + time[2]*10 + time[3]
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var d = make([]string, n)
	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d[i])
		tmp := strings.Split(d[i], "-")

		a[i] = toSecondFirst(strToDigits(tmp[0]))
		b[i] = toSecondSecond(strToDigits(tmp[1]))
	}

	t := make([]int, 24*60*60+1)
	for i := 0; i < n; i++ {
		t[a[i]]++
		t[b[i]+1]--
	}

	for i := 0; i < 24*60*60; i++ {
		t[i+1] += t[i]
	}

	flag := false
	ans := []int{}
	for i := 0; i < 24*60*60; i++ {
		if t[i] <= 0 && flag {
			ans = append(ans, i-1)
			flag = false
		} else if t[i] <= 0 {
			flag = false
			continue
		}

		if t[i] > 0 && flag {

		} else if t[i] > 0 {
			flag = true
			ans = append(ans, i)
		}

	}

	for i := 0; i < len(ans); i += 2 {
		if ans[i]/60 < 10 {
			fmt.Fprintf(writer, "0%v", ans[i]/60)
		} else {
			fmt.Fprintf(writer, "%v", ans[i]/60)
		}
		if ans[i]%60 < 10 {
			fmt.Fprintf(writer, "0%v", ans[i]%60)
		} else {
			fmt.Fprintf(writer, "%v", ans[i]%60)
		}
		fmt.Fprintf(writer, "-")
		if ans[i+1]/60 < 10 {
			fmt.Fprintf(writer, "0%v", ans[i+1]/60)
		} else {
			fmt.Fprintf(writer, "%v", ans[i+1]/60)
		}
		if ans[i+1]%60 < 10 {
			fmt.Fprintf(writer, "0%v", ans[i+1]%60)
		} else {
			fmt.Fprintf(writer, "%v", ans[i+1]%60)
		}
		fmt.Fprintf(writer, "\n")
	}

}
