package algo

type rect struct {
	h int
	w int
}

func maxAreaInHistgram(hist []int) (mx, w int) {
	h := []int{}
	for _, x := range hist {
		h = append(h, x)
	}
	h = append(h, 0)
	st := stack{}

	for i := range make([]int, len(h)) {
		if st.empty() {
			st.push(rect{h[i], i})
		} else if st.last().h <= h[i] {
			st.push(rect{h[i], i})
		} else if st.last().h > h[i] {
			j := i
			for !st.empty() && st.last().h > h[i] {
				rect := st.pop()
				j = rect.w
				mx = max(mx, (i-j)*rect.h)
				if mx == (i-j)*rect.h {
					w = i - j
				}
			}
			st.push(rect{h[i], j})
		}
	}
	return
}
