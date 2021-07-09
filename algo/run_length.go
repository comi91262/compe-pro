package algo

type runLength struct {
	s string
	l int
}

func encodeRunLength(s []string) []runLength {
	r := make([]runLength, len(s))

	cnt := 0
	for i := 0; i < len(s); i++ {
		if cnt > 0 && r[cnt-1].s == s[i] {
			r[cnt-1].l++
		} else {
			r[cnt].s = s[i]
			r[cnt].l = 1
			cnt++
		}
	}

	return r
}
