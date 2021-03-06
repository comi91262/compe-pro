package algo

type runLength struct {
	s string
	l int
}

func encodeRunLength(s []string) (r []runLength) {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if 0 < cnt && r[cnt-1].s == s[i] {
			r[cnt-1].l++
		} else {
			r = append(r, runLength{s: s[i], l: 1})
			cnt++
		}
	}
	return
}
